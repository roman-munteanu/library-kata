package models

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	TakenByUser string `json:"taken_by_user"`
}

type Request struct {
	UserID string `json:"user_id"`
	BookID string `json:"book_id"`
}

type Response struct {
	IsSuccess bool `json:"is_success"`
}
