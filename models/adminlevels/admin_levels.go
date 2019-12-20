package adminlevels

type (

	// AdminLevelModel .
	AdminLevelModel struct {
		Name   string `json:"name" validate:"required" example:"test"`
		Level  int    `json:"level" validate:"required" example:"2"`
		Enable *int   `json:"enable"`
	}

	// AdminLevelOption .
	AdminLevelOption struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
