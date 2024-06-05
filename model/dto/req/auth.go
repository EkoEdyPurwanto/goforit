package req

type (
	AuthRegisterRequest struct {
		Username        string `json:"username" validate:"required,min=3,max=15"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,min=6,max=20"`
		PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	}
)
