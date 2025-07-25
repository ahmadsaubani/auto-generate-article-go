package helpers

import (
	"fmt"
	"news-go/src/configs/database"
	"news-go/src/traits"
	"news-go/src/utils/filters"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const maxBatchSize = 500
const layout = "2006-01-02"

func StringToDateOnly(input string) time.Time {
	t, _ := time.Parse(layout, input)
	return t
}

func GeneratePoints(price int) uint {
	if price <= 0 {
		return 0
	}
	return uint(price / 1000)
}

func InsertModelBatch[T any](models []T) error {
	if len(models) == 0 {
		return nil
	}

	if database.GormDB == nil {
		return fmt.Errorf("❌ Database connection is not initialized")
	}

	for start := 0; start < len(models); start += maxBatchSize {
		end := start + maxBatchSize
		if end > len(models) {
			end = len(models)
		}
		batch := models[start:end]

		// Set UUID untuk setiap item dalam batch
		for i := range batch {
			err := traits.SetUUIDForStruct(&batch[i])
			if err != nil {
				return fmt.Errorf("❌ Error setting UUID: %w", err)
			}
		}

		err := database.GormDB.Transaction(func(tx *gorm.DB) error {

			if err := tx.Create(&batch).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf("❌ GORM transaction failed: %w", err)
		}
	}

	return nil
}

func InsertModel[T any](model *T) error {
	if err := traits.SetUUIDForStruct(model); err != nil {
		return fmt.Errorf("❌ Error setting UUID: %w", err)
	}
	return database.GormDB.Create(model).Error
}

func GetAllModelsWithDB[T any](ctx *gin.Context, db *gorm.DB, models *[]T, pag PaginationParams) error {
	limit := pag.Limit
	offset := pag.Offset
	orderBy := ctx.DefaultQuery("order_by", "")

	if orderBy != "" {
		orderParts := strings.Split(orderBy, ",")
		if len(orderParts) == 2 {
			orderBy = fmt.Sprintf("%s %s", orderParts[0], orderParts[1])
		}
	}

	if db == nil {
		return fmt.Errorf("database connection not initialized")
	}

	query := db
	query = query.Limit(limit).Offset(offset)

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	whereClause, args, err := filters.BuildFilters(ctx)
	if err != nil {
		return err
	}
	if whereClause != "" {
		query = query.Where(whereClause, args...)
	}
	return query.Find(models).Error
}

func GetAllModels[T any](ctx *gin.Context, models *[]T) error {
	orderBy := ctx.DefaultQuery("order_by", "")

	if orderBy != "" {
		orderParts := strings.Split(orderBy, ",")
		if len(orderParts) == 2 {
			orderBy = fmt.Sprintf("%s %s", orderParts[0], orderParts[1])
		}
	}

	if database.GormDB == nil {
		return fmt.Errorf("database connection not initialized")
	}

	query := database.GormDB

	if limitStr := ctx.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			query = query.Limit(limit)
		}
	}

	if offsetStr := ctx.Query("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil && offset >= 0 {
			query = query.Offset(offset)
		}
	}

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	whereClause, args, err := filters.BuildFilters(ctx)
	if err != nil {
		return err
	}
	if whereClause != "" {
		query = query.Where(whereClause, args...)
	}

	return query.Find(models).Error
}

func GetModelByID[T any](model *T, id any) error {
	return database.GormDB.First(model, id).Error
}

func UpdateModelByIDWithMap[T any](updatedFields map[string]interface{}, id any) error {
	return database.GormDB.Model(new(T)).Where("id = ?", id).Updates(updatedFields).Error
}

func UpdateModelByID[T any](model *T, id any) error {
	return database.GormDB.Model(model).Where("id = ?", id).Updates(model).Error

}

func DeleteModelByID[T any](model *T, id any) error {
	return database.GormDB.Delete(model, id).Error
}

func FindOneByField[T any](model *T, conditions ...any) error {
	return FindOneByFieldWithPreload(model, nil, conditions...)
}

func FindOneByFieldWithPreload[T any](model *T, preloads []string, conditions ...any) error {
	if len(conditions)%2 != 0 {
		return fmt.Errorf("conditions must be in key-value pairs")
	}

	query := database.GormDB
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	for i := 0; i < len(conditions); i += 2 {
		field := conditions[i].(string)
		value := conditions[i+1]
		query = query.Where(fmt.Sprintf("%s = ?", field), value)
	}

	return query.First(model).Error
}
