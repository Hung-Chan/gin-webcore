package auth

type (
	// Login 登入者帳密.
	// Account 限制長度最小為四個字元，最大20個字元，且必填
	// Password 限制長度最小為四個字元，最大20個字元，且必填
	Login struct {
		Account  string `json:"account" validate:"required,min=4,max=20" example:"admin"`
		Password string `json:"password" validate:"required,min=4,max=20" example:"qaz123"`
	}
)
