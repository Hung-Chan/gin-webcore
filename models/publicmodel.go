package models

import "time"

// QueryModel : 此Model為所有查詢功能固定所帶參數 .
type QueryModel struct {
	Page          int    `form:"page"`
	Limit         int    `form:"limit"`
	SortColumn    string `form:"sortColumn"`
	SortDirection string `form:"sortDirection"`
	Enable        int    `form:"enable"`
	Name          string `form:"name"`
}

// IDInfo .
type IDInfo struct {
	ID        *int      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
