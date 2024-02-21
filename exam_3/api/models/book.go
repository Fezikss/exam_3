package models

type Book struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AuthorName string `json:"author_name"`
	PageNumber int    `json:"page_number"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateBook struct {
	Name       string `json:"name"`
	AuthorName string `json:"author_name"`
	PageNumber int    `json:"page_number"`
}

type UpdateBook struct {
	ID         string `json:"-"`
	Name       string `json:"name"`
	AuthorName string `json:"author_name"`
}

type UpdateBookPageNumber struct {
	ID         string `json:"-"`
	PageNumber int    `json:"page_number"`
}

type BooksResponse struct {
	Books []Book `json:"books"`
	Count int    `json:"count"`
}
