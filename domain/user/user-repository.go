package user

type UserRepository interface {
	RegisterUser(payload *CreateUser) (*UserLoginResponse, error)
	UserLogin(payload *UserLogin) (*UserLoginResponse, error)
	GetUserById(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByPhone(phoneNumbers string) (*User, error)
}
