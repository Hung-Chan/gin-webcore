package model

import (
	"time"
)

// Model .
type Model struct {
	ID        int       `gorm:"type: int unsigned auto_increment; primary_key"`
	CreatedAt time.Time `gorm:"comment='建立時間'"`
	UpdatedAt time.Time `gorm:"comment='最後更新時間'"`
}
