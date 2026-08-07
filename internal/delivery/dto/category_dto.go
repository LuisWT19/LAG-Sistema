package dto

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description`
	Slug        string `json:"slug"`
	IsActive    bool   `json:"is_active"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description`
	Slug        string `json:"slug"`
	IsActive    bool   `json:"is_active"`
}
