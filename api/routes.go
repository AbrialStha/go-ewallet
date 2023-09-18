package api

import (
	"net/http"

	"github.com/abrialstha/go-ewallet/utils"
	"github.com/gorilla/mux"
)

func RegisterUserRouter(router *mux.Router) {
	router.HandleFunc("/me", utils.ApiHandler(GetUser)).Methods("GET")
	router.HandleFunc("/signup", utils.ApiHandler(GetUser)).Methods("POST")
	router.HandleFunc("/login", utils.ApiHandler(GetUser)).Methods("POST")
	router.HandleFunc("/logout", utils.ApiHandler(GetUser)).Methods("GET")
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
