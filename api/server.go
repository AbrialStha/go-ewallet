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

	//? Initialize the DB (migration stuff will go here)
	client.InitPgDB()

	//? Register the Routes here
	RegisterUserRouter(router)

	log.Printf("Server is listening on port %s...\n", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Println("Error While Starting Server:", err)
	}
}
