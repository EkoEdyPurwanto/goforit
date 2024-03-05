package req

type (
	RegisterUsersRequest struct {
		Identifier      interface{} `json:"identifier"`
		Username        string      `json:"username"`
		Password        string      `json:"password"`
		PasswordConfirm string      `json:"password_confirm"`
	}
)
