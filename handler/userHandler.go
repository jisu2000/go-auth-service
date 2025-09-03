package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jisu2000/go-auth-service/dto/request"
	"github.com/jisu2000/go-auth-service/service"
	"github.com/jisu2000/go-auth-service/wrapper"
)

type UserHandler struct {
	SRV *service.UserService
}

func (srv *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	userRegisterRequest := request.UserRegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&userRegisterRequest)
	if err != nil {
		wrapper.JSON(w, http.StatusBadRequest, false, nil, "Invalid request body")
		return
	}
	res, err := srv.SRV.CreateNewUser(userRegisterRequest)
	if err != nil {
		wrapper.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	wrapper.JSON(w, http.StatusOK, true, res, "")

}
