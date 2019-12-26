package admingroups

import (
	"encoding/json"
)

type (
	// AdminGroupModel .
	AdminGroupModel struct {
		Name       string          `json:"name" validate:"required" example:"test"`
		Permission json.RawMessage `json:"permission"`
		Remark     string          `json:"remark" example:"test"`
		Enable     *int            `json:"enable"`
	}

	// AdminGroupOption .
	AdminGroupOption struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	// Permission .
	Permission struct {
		Permission json.RawMessage `json:"permission"`
	}
)
