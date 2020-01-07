package administrators

type (
	// AdministratorModel .
	AdministratorModel struct {
		Name     string `json:"name" validate:"required" example:"test"`
		Account  string `json:"account" validate:"required" example:"test"`
		Password string `json:"password" example:"qaz123"`
		GroupID  *int   `json:"group_id" validate:"required" example:"1"`
		LevelID  *int   `json:"level_id" validate:"required" example:"1"`
		Enable   *int   `json:"enable"`
		Remark   string `json:"remark"`
		AdminID  int    `json:"admin_id"`
	}

	// Administrator .
	Administrator struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
