package admingroups

import (
	"encoding/json"
)

type (
	// AdminGroup .
	AdminGroup struct {
		Name       string          `json:"name" validate:"required"`
		Permission json.RawMessage `json:"permission"`
		Remark     string          `json:"remark"`
		Enable     *int            `json:"enable"`
	}

	// AdminGroupOption .
	AdminGroupOption struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
