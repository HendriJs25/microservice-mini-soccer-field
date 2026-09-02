package request

type RegisterRequest struct {
	Name            string `json:"name" validate:"required,notblank,max=100"`
	Username        string `json:"username" validate:"required,notblank,max=20"`
	Password        string `json:"password" validate:"required,notblank,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8,eqfield=Password"`
	Email           string `json:"email" validate:"required,notblank,email"`
}
