package usecase

import (
	"log"
	"time"

	"github.com/abrialstha/go-ewallet/domain/user"
)

type userUseCase struct {
	userRepository user.UserRepository
}

func NewUserUsecase(repo *user.UserRepository) *userUseCase {
	return &userUseCase{
		userRepository: *repo,
	}
}

func (i *userUseCase) RegisterUser(payload *user.CreateUser) (*user.UserLoginResponse, error) {
	// Update Created
	payload.CreatedAt = time.Now().UTC()
	repo := i.userRepository
	user, err := repo.RegisterUser(payload)
	if err == nil {
		log.Print("Sending Email to user for sucessful registration")
	}
	return user, err
}
