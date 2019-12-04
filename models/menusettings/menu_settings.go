package menusettings

type (
	// Menusetting .
	Menusetting struct {
		ParentID int    `json:"parent_id"`
		Code     string `json:"code"`
		Name     string `json:"name"`
		GroupID  int    `json:"group_id"`
		Icon     string `json:"icon"`
		Icolor   string `json:"icolor"`
		Access   string `json:"access"`
		Sort     int    `json:"sort"`
		Enable   int    `json:"enable"`
	}
)
