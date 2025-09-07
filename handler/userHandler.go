package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jisu2000/go-auth-service/dto/request"
	"github.com/jisu2000/go-auth-service/dto/response"
	"github.com/jisu2000/go-auth-service/jwt"
	"github.com/jisu2000/go-auth-service/service"
	"github.com/jisu2000/go-auth-service/util"
	"github.com/jisu2000/go-auth-service/wrapper"
)

type UserHandler struct {
	SRV  *service.UserService
	RSRV *service.RefreshTokenService
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

	jwtToken, _, err := jwt.CreateToken(res.Email, res.Roles)

	if err != nil {
		wrapper.JSON(w, http.StatusInternalServerError, false, nil, "Error generating token")
		return
	}

	refreshToken, err := srv.RSRV.CreateNewRefreshToken(res.Id)
	if err != nil {
		wrapper.JSON(w, http.StatusInternalServerError, false, nil, "Error generating refresh token")
		return
	}

	authRes := response.AuthResponse{
		AccessToken: jwtToken,
		RefreshToken: func() string {
			if refreshToken == nil {
				return ""
			}
			return *refreshToken
		}(),
	}

	wrapper.JSON(w, http.StatusOK, true, authRes, "")

}

func (srv *UserHandler) FetchUserList(w http.ResponseWriter, r *http.Request) {
	userList, err := srv.SRV.FetchAllUsers()
	if err != nil {
		wrapper.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	wrapper.JSON(w, http.StatusOK, true, userList, "")
}

func (srv *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		wrapper.JSON(w, http.StatusBadRequest, false, nil, "invalid user id")
		return
	}
	if err := srv.SRV.DeleteUser(id); err != nil {
		if err.Error() == "user not found" {
			wrapper.JSON(w, http.StatusNotFound, false, nil, err.Error())
			return
		}
		wrapper.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	wrapper.JSON(w, http.StatusOK, true, "User Deleted", "")
}

func (srv *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	loginReq := request.UserLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		wrapper.JSON(w, http.StatusBadRequest, false, nil, "Invalid request body")
		return
	}
	res, err := srv.SRV.FetchUserFromEmailOrMobile(loginReq.Username)

	if err != nil {
		wrapper.JSON(w, http.StatusBadRequest, false, nil, "Invalid Email or Mobile Number")
	}

	// handle nil user (not found)
	if res == nil {
		wrapper.JSON(w, http.StatusBadRequest, false, nil, "Invalid Email or Mobile Number")
		return
	}
	// check password first
	if passwordMatched := util.CheckPassword(res.Password, loginReq.Password); !passwordMatched {
		wrapper.JSON(w, http.StatusBadRequest, false, nil, "Invalid Password")
		return
	}
	// only generate token after successful password verification
	jwtToken, _, err := jwt.CreateToken(res.Email, res.Roles)
	if err != nil {
		wrapper.JSON(w, http.StatusInternalServerError, false, nil, "Error generating token")
		return
	}

	refreshToken, err := srv.RSRV.CreateNewRefreshToken(res.Id)
	if err != nil {
		wrapper.JSON(w, http.StatusInternalServerError, false, nil, "Error generating refresh token")
		return
	}

	authRes := response.AuthResponse{
		AccessToken:  jwtToken,
		RefreshToken: *refreshToken,
	}
	wrapper.JSON(w, http.StatusOK, true, authRes, "")

}

func (srv *UserHandler) VerifyAndGenerateRefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshTokenRequest := request.TokenRefreshRequest{}

	if err := json.NewDecoder(r.Body).Decode(&refreshTokenRequest); err != nil {
		wrapper.JSON(w, http.StatusBadRequest, false, nil, "Invalid request body")
		return
	}

	refreshToken, userId, err := srv.RSRV.VerifyRefreshToken(refreshTokenRequest.RefreshToken)
	if err != nil || userId == nil || refreshToken == nil {
		wrapper.JSON(w, http.StatusBadRequest, false, nil, "Invalid refresh token")
		return
	}

	userRes, err := srv.SRV.GetUserById(*userId)
	if err != nil {
		wrapper.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	jwtToken, _, err := jwt.CreateToken(userRes.Email, userRes.Roles)
	if err != nil {
		wrapper.JSON(w, http.StatusInternalServerError, false, nil, "Error generating token")
	}
	wrapper.JSON(w, http.StatusOK, true, response.AuthResponse{AccessToken: jwtToken, RefreshToken: *refreshToken}, "")
}
