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
	roleRepo := &repo.RoleRepo{DB: database}
	refreshTokenRepo := &repo.RefreshTokenRepo{DB: database}
	userService := &service.UserService{REPO: userRepo, ROLE_REPO: roleRepo}
	referTokenService := &service.RefreshTokenService{USERREPO: userRepo, REFRESHREPO: refreshTokenRepo}
	userHandler := &handler.UserHandler{SRV: userService, RSRV: referTokenService}
	r := mux.NewRouter()
	r.HandleFunc("/users/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", userHandler.LoginUser).Methods("POST")
	r.HandleFunc("/users/refresh_token", userHandler.VerifyAndGenerateRefreshToken).Methods("POST")
	r.HandleFunc("/users/fetch_all", userHandler.FetchUserList).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
