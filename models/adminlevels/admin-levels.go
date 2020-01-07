package adminlevels

type (

	// AdminLevelModel .
	// Name: 必填、長度20字元
	// Level: 必填
	// Enable: 資料庫預設 0
	AdminLevelModel struct {
		Name   string `json:"name" validate:"required,max=20" example:"test"`
		Level  int    `json:"level" validate:"required" example:"2"`
		Enable *int   `json:"enable"`
	}

	// AdminLevelOption .
	AdminLevelOption struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
