package models

import "time"

// PublicQueryModel : 此Model為所有查詢功能固定所帶參數 .
type PublicQueryModel struct {
	Page          *int    `form:"page"`
	Limit         *int    `form:"limit"`
	SortColumn    *string `form:"sortColumn"`
	SortDirection *string `form:"sortDirection"`
}

// IDInfo .
type IDInfo struct {
	ID        *int      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
