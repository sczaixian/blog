package request

type CreateArticle struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Excerpt    string `json:"excerpt"`
	CategoryID uint   `json:"category_id"`
}

type EditArticle struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Excerpt    string `json:"excerpt"`
	CategoryID uint   `json:"category_id"`
}

type SaveArticle struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Excerpt    string `json:"excerpt"`
	CategoryID uint   `json:"category_id"`
}
