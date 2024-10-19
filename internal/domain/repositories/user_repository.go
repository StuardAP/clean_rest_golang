package repositories


import (
    entity "github.com/StuardAP/clean_rest_golang/internal/domain/entities"
)

type UserRepository interface {
    CreateUser(user *entity.User) error
    GetUserByEmail(email string) (*entity.User, error)
    GetAllUsers() ([]entity.User, error)
    UpdateUser(user *entity.User) error
    DeleteUser(id string) error
}
