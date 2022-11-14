package user_handler_rest

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
