package administrators

import (
	"encoding/json"
)

type (
	// Administrator .
	Administrator struct {
		Name     string `json:"name"`
		Account  string `json:"account"`
		Password string `json:"password"`
		GroupID  *int   `json:"group_id"`
		LevelID  *int   `json:"level_id"`
		Enable   *int   `json:"enable"`
		Remark   string `json:"remark"`
	}

	// NewPermission .
	NewPermission struct {
		Permission json.RawMessage `json:"permission"`
	}
)
