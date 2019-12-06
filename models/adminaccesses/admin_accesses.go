package adminaccesses

// AdminAccess .
type AdminAccess struct {
	Name   string `json:"name" validate:"required"`
	Code   string `json:"code" validate:"required"`
	Enable *int   `json:"enable"`
}
