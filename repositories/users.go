package repositories

import (
	"github.com/google/uuid"

	"roman-munteanu/library-kata/models"
)

var usersData []models.User

type UsersRepositoryAPI interface {
	FetchAll() ([]models.User, error)
	Save(u models.User) (bool, error)
}

type UsersRepository struct {
}

func NewUsersRepository() *UsersRepository {
	usersData = []models.User{
		{
			ID:   "cb0aa0ba-3c04-4ee4-9c2c-a34bd3dbc7de",
			Name: "User 1",
		},
		{
			ID:   "a122c28c-790a-4470-bd55-880674afdce9",
			Name: "User 2",
		},
	}

	return &UsersRepository{}
}

func (r *UsersRepository) FetchAll() ([]models.User, error) {
	return usersData, nil
}

func (r *UsersRepository) Save(u models.User) (bool, error) {
	u.ID = uuid.New().String()
	usersData = append(usersData, u)

	return true, nil
}
