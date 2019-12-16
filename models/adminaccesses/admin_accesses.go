package adminaccesses

type (
	// AdminAccess .
	AdminAccess struct {
		Name   string `json:"name" validate:"required"`
		Code   string `json:"code" validate:"required"`
		Enable *int   `json:"enable"`
	}
)
