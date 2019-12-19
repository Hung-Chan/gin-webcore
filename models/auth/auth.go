package auth

type (
	// Login .
	Login struct {
		Account  string `json:"account" validate:"required,min=4,max=20" example:"admin"`
		Password string `json:"password" validate:"required,min=4,max=20" example:"qaz123"`
	}
)
