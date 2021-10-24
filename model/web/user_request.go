package web

type UserRequest struct {
	Email    string `validate:"required" json:"email"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
