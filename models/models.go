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

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

type GenericError struct {
	Message string
}

func (e GenericError) Error() string {
	return e.Message
}

type Response struct {
	IsSuccess bool `json:"is_success"`
}
