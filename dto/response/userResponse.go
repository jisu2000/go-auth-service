package response

type UserResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	CreatedAt    string `json:"created_at"`
}
