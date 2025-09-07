package request

type UserRegisterRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MobileNumber string `json:"mobile_number"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
