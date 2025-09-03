package response

type UserResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	CreatedAt    string `json:"created_at"`
}
