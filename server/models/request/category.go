package request

type CreateCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
