package repositories

import (
	"roman-munteanu/library-kata/models"
)

type UsersRepositoryAPI interface {
	FetchAll() ([]models.User, error)
}

type UsersRepository struct {
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) FetchAll() ([]models.User, error) {
	data := []models.User{
		{
			ID:   "uuid1",
			Name: "name 1",
		},
		{
			ID:   "uuid2",
			Name: "name 2",
		},
	}

	return data, nil
}
