package adminaccesses

type (
	// AdminAccessModel .
	AdminAccessModel struct {
		Name   string `json:"name" validate:"required" example:"test"`
		Code   string `json:"code" validate:"required" example:"test"`
		Enable *int   `json:"enable"`
	}
)
