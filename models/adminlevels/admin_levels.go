package adminlevels

type (

	// AdminLevel .
	AdminLevel struct {
		Name   string `json:"name" validate:"required"`
		Level  int    `json:"level" validate:"required"`
		Enable int    `json:"enable"`
	}

	// AdminLevelOption .
	AdminLevelOption struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
