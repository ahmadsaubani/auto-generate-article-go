package category_dtos

type CategoryDTO struct {
	ID    uint   `json:"id"`
	UUID  string `json:"uuid"`
	Label string `json:"label"`
	Value string `json:"value"`
}
