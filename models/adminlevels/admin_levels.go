package adminlevels

type (

	// AdminLevel .
	AdminLevel struct {
		Name   string `json:"name"`
		Level  int    `json:"level"`
		Enable int    `json:"enable"`
	}
)
