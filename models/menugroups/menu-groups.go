package menugroups

type (
	// MenuGroupModel .
	MenuGroupModel struct {
		Name   string `json:"name" validate:"required" example:"test"`
		Enable *int   `json:"enable"`
		Sort   int    `json:"sort"`
	}

	// MenuGroupOption .
	MenuGroupOption struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
