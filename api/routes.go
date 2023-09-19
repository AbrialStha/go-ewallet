package api

import (
	"net/http"

	domain "github.com/abrialstha/go-ewallet/domain/user"
	infrastructure "github.com/abrialstha/go-ewallet/infrastructure/user"
	usecase "github.com/abrialstha/go-ewallet/usercase/user"
	"github.com/abrialstha/go-ewallet/utils"
	"github.com/gorilla/mux"
)

func RegisterUserRouter(router *mux.Router) {
	router.HandleFunc("/signup", utils.ApiHandler(registerUser)).Methods("POST")
	router.HandleFunc("/login", utils.ApiHandler(GetUser)).Methods("POST")
	router.HandleFunc("/logout", utils.ApiHandler(GetUser)).Methods("GET")
	router.HandleFunc("/me", utils.ApiHandler(GetUser)).Methods("GET")
}

// This will be calling the application usecase most of the time
func registerUser(w http.ResponseWriter, r *http.Request) error {

	payload := &domain.CreateUser{}
	utils.ParseBody(r, payload)

	// Init userUseCase
	userRepo := infrastructure.NewUserRepository()
	user := usecase.NewUserUsecase(&userRepo)

	res, err := user.RegisterUser(payload)
	if err != nil {
		return err
	}

	return utils.JSON(w, http.StatusOK, res)
}

// This will be calling the application usecase most of the time
func GetUser(w http.ResponseWriter, r *http.Request) error {
	type HelloReturn struct {
		Message string `json:"message"`
	}

	c := HelloReturn{
		Message: "Hello World",
	}

	return utils.JSON(w, http.StatusOK, c)
}
