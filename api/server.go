package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abrialstha/go-ewallet/client"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr int) *Server {
	return &Server{
		listenAddr: fmt.Sprintf(":%d", listenAddr),
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()

	//? Connect Db
	client.PgConnect()

	//? For Testing Purpose, we will call this in the repository
	db := client.GetPgDB()
	fmt.Printf("%+v \n", db)

	//? Register the Routes here
	RegisterUserRouter(router)

	log.Printf("Server is listening on port %s...\n", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Println("Error While Starting Server:", err)
	}
}
