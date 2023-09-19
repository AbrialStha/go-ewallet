package infrastructure

import (
	"database/sql"
	"log"

	"github.com/abrialstha/go-ewallet/client"
	"github.com/abrialstha/go-ewallet/domain/user"
)

type userRepositoryImpl struct {
	dbClient *sql.DB
}

func (i *userRepositoryImpl) RegisterUser(payload *user.CreateUser) (*user.UserLoginResponse, error) {
	log.Printf("%+v", payload)

	query := `
	insert into public.user
	(first_name, middle_name, last_name, email, phone_number, password, created_at)
	values
	($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := i.dbClient.Query(
		query,
		payload.FirstName,
		payload.MiddleName,
		payload.LastName,
		payload.Email,
		payload.PhoneNumber,
		payload.Password,
		payload.CreatedAt,
	)

	if err != nil {
		log.Printf("Error: %+v", err)
		return nil, err
	}

	//? Login Action after the signup will be here
	resp := new(user.UserLoginResponse)
	resp.Type = "Bearer"
	resp.Token = "jwt_token_will_be_here"

	return resp, nil
}

func (i *userRepositoryImpl) UserLogin(payload *user.UserLogin) (*user.UserLoginResponse, error) {
	log.Printf("%+v", payload)
	res := new(user.UserLoginResponse)
	return res, nil
}

func (i *userRepositoryImpl) GetUserById(id string) (*user.User, error) {
	log.Printf("%+v", id)
	user := new(user.User)
	return user, nil
}

func (i *userRepositoryImpl) GetUserByEmail(email string) (*user.User, error) {
	log.Printf("%s", email)
	user := new(user.User)
	return user, nil
}

func (i *userRepositoryImpl) GetUserByPhone(phoneNumbers string) (*user.User, error) {
	log.Printf("%s", phoneNumbers)
	user := new(user.User)
	return user, nil
}

func NewUserRepository() user.UserRepository {
	db := client.GetPgDB()
	return &userRepositoryImpl{
		dbClient: db,
	}
}
