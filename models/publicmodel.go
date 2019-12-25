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

// QueryModelNew : 此Model為所有查詢功能固定所帶參數 .
type QueryModelNew struct {
	Page          int     `form:"page" example:"1"`
	Limit         int     `form:"limit" example:"10"`
	SortColumn    string  `form:"sortColumn" example:"id"`
	SortDirection string  `form:"sortDirection" example:"asc"`
	Enable        *int    `form:"enable" example:"1"`
	Name          *string `form:"name" example:""`
	IP            *string `form:"ip" example:""`
	Subnet        *string `form:"subnet" example:""`
	Country       *string `form:"country" example:""`
	Group         *int    `form:"group" example:""`
	Level         *int    `form:"level" example:""`
	NameItem      *string `form:"nameItem" example:"account"`
	AccountOrName *string `form:"accountOrName" example:""`
}

// IDInfo .
type IDInfo struct {
	ID        *int      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewQueryModel .
func NewQueryModel() *QueryModelNew {
	return &QueryModelNew{
		Page:          1,
		Limit:         10,
		SortColumn:    "id",
		SortDirection: "asc",
		Enable:        nil,
		Name:          nil,
		IP:            nil,
		Subnet:        nil,
		Country:       nil,
		Group:         nil,
		Level:         nil,
		NameItem:      nil,
		AccountOrName: nil,
	}
}
