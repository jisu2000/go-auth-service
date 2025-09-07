package response

type UserResponse struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	MobileNumber string   `json:"mobile_number"`
	CreatedAt    string   `json:"created_at"`
	Password     string   `json:"-"`
	Roles        []string `json:"roles"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
