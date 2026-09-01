package request

type RegisterRequest struct {
	Name            string `json:"name" validate:"required, max=100"`
	Username        string `json:"username" validate:"required,max=20"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8,eqfield=Password"`
	Email           string `json:"email" validate:"required,email"`
}
