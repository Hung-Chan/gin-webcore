package menugroups

type (
	// MenuGroupModel .
	MenuGroupModel struct {
		Name   string `json:"name" validate:"required" example:"test"`
		Enable *int   `json:"enable"`
		Sort   int    `json:"sort"`
	}
)
