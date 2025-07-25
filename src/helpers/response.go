package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"news-go/src/configs/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PaginationParams struct {
	Page   int
	Limit  int
	Offset int
}

type PaginationMeta struct {
	Page  int    `json:"page"`
	Limit int    `json:"per_page"`
	Total int64  `json:"total"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
}

type Response struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Links   interface{} `json:"links,omitempty"`
}

type ErrorData struct {
	Error       interface{} `json:"error"`
	Path        string      `json:"path"`
	Method      string      `json:"method"`
	ClientIP    string      `json:"clientIP"`
	Status      int         `json:"status"`
	RequestBody interface{} `json:"requestBody"`
}

func SuccessResponse(ctx *gin.Context, message string, data interface{}, pagination ...PaginationMeta) {
	if message == "" {
		message = "Data Found"
	}

	webResponse := Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	if len(pagination) > 0 {
		webResponse.Meta = map[string]interface{}{
			"pagination": pagination[0],
		}
		webResponse.Links = buildPaginationLinks(ctx, pagination[0])
	}

	JSONResponse(ctx, webResponse)
}

func GetPaginatedData[T any](ctx *gin.Context, db *gorm.DB, order string, page, limit, offset int) ([]T, PaginationMeta, int64) {
	var data []T
	var total int64

	db.Model(new(T)).Count(&total)

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	if page > totalPages && totalPages != 0 {
		page = totalPages
	}

	if order == "" {
		order = "created_at desc"
	}

	db.Limit(limit).
		Offset(offset).
		Order(order).
		Find(&data)

	meta := PaginationMeta{
		Page:  page,
		Limit: limit,
		Total: total,
	}

	return data, meta, total
}

func PaginatedResponse(ctx *gin.Context, message string, data interface{}, page, limit int, total int64) {
	meta := PaginationMeta{
		Page:  page,
		Limit: limit,
		Total: total,
	}

	SuccessResponse(ctx, message, data, meta)
}

func buildPaginationLinks(ctx *gin.Context, meta PaginationMeta) map[string]string {
	links := make(map[string]string)

	// Calculate total pages
	totalPages := int(math.Ceil(float64(meta.Total) / float64(meta.Limit)))

	if meta.Page < totalPages {
		links["next"] = buildPaginationLink(ctx, meta.Page+1, meta.Limit)
	}
	if meta.Page > 1 {
		links["prev"] = buildPaginationLink(ctx, meta.Page-1, meta.Limit)
	}
	if meta.Page > 1 {
		links["first"] = buildPaginationLink(ctx, 1, meta.Limit)
	}
	if meta.Page < totalPages {
		links["last"] = buildPaginationLink(ctx, totalPages, meta.Limit)
	}

	return links
}

func buildPaginationLink(ctx *gin.Context, page, limit int) string {
	return fmt.Sprintf("%s?page=%d&per_page=%d", ctx.Request.URL.Path, page, limit)
}

func ErrorResponse(ctx *gin.Context, err error, httpCode ...int) {
	if len(httpCode) == 0 {
		httpCode = append(httpCode, http.StatusBadRequest)
	}

	message := ParseValidationError(err)

	webResponse := Response{
		Success: false,
		Message: message,
		Data:    nil,
	}
	ctx.JSON(httpCode[0], webResponse)
}

func ParseValidationError(err error) interface{} {
	var ve validator.ValidationErrors
	var message interface{}

	if errors.As(err, &ve) {
		errorMap := map[string]string{}
		for _, fe := range ve {
			field := strings.ToLower(fe.Field())
			switch fe.Tag() {
			case "required":
				errorMap[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errorMap[field] = fmt.Sprintf("%s must be a valid email", field)
			case "min":
				errorMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fe.Param())
			default:
				errorMap[field] = fmt.Sprintf("%s is invalid", field)
			}
		}
		jsonMessage, _ := json.Marshal(errorMap)
		message = string(jsonMessage)
		return message
	}
	return err.Error()
}

func JSONResponse(ctx *gin.Context, data interface{}) {
	isCreate := ctx.Request.Method == http.MethodPost
	statusCode := http.StatusOK
	if isCreate {
		statusCode = http.StatusCreated
	}

	ctx.JSON(statusCode, data)
}

// func GetPaginationParams(ctx *gin.Context) (page, limit, offset int) {
// 	page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))
// 	limit, _ = strconv.Atoi(ctx.DefaultQuery("per_page", "10"))

// 	if page < 1 {
// 		page = 1
// 	}
// 	if limit < 1 || limit > 100 {
// 		limit = 10
// 	}

// 	offset = (page - 1) * limit
// 	return
// }

func GetPaginationParams(ctx *gin.Context) PaginationParams {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	return PaginationParams{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}

func CountModel[T any]() (int64, error) {
	var total int64
	err := database.GormDB.Model(new(T)).Count(&total).Error
	return total, err
}
