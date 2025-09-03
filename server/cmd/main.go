package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jisu2000/go-auth-service/db"
	"github.com/jisu2000/go-auth-service/handler"
	"github.com/jisu2000/go-auth-service/repo"
	"github.com/jisu2000/go-auth-service/service"
)

func main() {
	database := db.Connect()
	userRepo := &repo.UserRepo{DB: database}
	userService := &service.UserService{REPO: userRepo}
	userHandler := &handler.UserHandler{SRV: userService}
	r := mux.NewRouter()
	r.HandleFunc("/users/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/users/fetch_all",userHandler.FetchUserList).Methods("GET")
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
