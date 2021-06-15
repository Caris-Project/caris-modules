package model

import (
	"time"

	"github.com/lib/pq"
)

// StaffRole ...
type StaffRole struct {
	ID          string
	Name        string
	Desc        string
	Role        string         `gorm:"primaryKey"`
	Permissions pq.StringArray `gorm:"type:text[]"`
	CreatedAt   time.Time
}

// TableName overrides the table name
func (StaffRole) TableName() string {
	return "staff_roles"
}
