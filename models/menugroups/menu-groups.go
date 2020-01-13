package menugroups

type (
	// MenuGroupModel .
	MenuGroupModel struct {
		Name   string `json:"name" validate:"required,max=20" example:"test"`
		Enable *int   `json:"enable"`
	}

	// MenuGroupOption .
	MenuGroupOption struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
