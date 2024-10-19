package usecases

import (
    entity "github.com/StuardAP/clean_rest_golang/internal/domain/entities"
    repository "github.com/StuardAP/clean_rest_golang/internal/domain/repositories"
    auth "github.com/StuardAP/clean_rest_golang/pkg/utilities/auth"
    "github.com/google/uuid"
    "time"
)

type UserUseCase struct {
    userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) *UserUseCase {
    return &UserUseCase{userRepository: userRepository}
}

func (u *UserUseCase) RegisterUser(name, email, password string) (*entity.User, error) {
    hashedPassword, err := auth.HashPassword(password)
    if err != nil {
        return nil, err
    }

    user := &entity.User{
        ID:        uuid.New(),
        Name:      name,
        Email:     email,
        Password:  hashedPassword,
        CreatedAt: time.Now(),
    }
    if err := u.userRepository.CreateUser(user); err != nil {
        return nil, err
    }
    return user, nil
}

func (u *UserUseCase) GetUserByEmail(email string) (*entity.User, error) {
		return u.userRepository.GetUserByEmail(email)
}

func (u *UserUseCase) GetAllUsers() ([]entity.User, error) {
        return u.userRepository.GetAllUsers()
}
