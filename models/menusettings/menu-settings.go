package menusettings

import (
	"encoding/json"
)

type (
	// Menusetting .
	Menusetting struct {
		ParentID int             `json:"parent_id"`
		Code     string          `json:"code"`
		Name     string          `json:"name"`
		GroupID  int             `json:"group_id"`
		Icon     string          `json:"icon"`
		Icolor   string          `json:"icolor"`
		Access   json.RawMessage `json:"access"`
		Sort     int             `json:"sort"`
		Enable   int             `json:"enable"`
	}

	// MenusettingModel .
	MenusettingModel struct {
		ParentID int             `json:"parent_id"`
		Code     string          `json:"code"`
		Name     string          `json:"name"`
		GroupID  int             `json:"group_id"`
		Icon     string          `json:"icon"`
		Icolor   string          `json:"icolor"`
		Access   json.RawMessage `json:"access"`
		Sort     int             `json:"sort"`
		Enable   int             `json:"enable"`
	}

	// Permission .
	Permission struct {
		ID     int             `json:"id"`
		Code   string          `json:"code"`
		Name   string          `json:"name"`
		Access json.RawMessage `json:"access"`
	}
)
