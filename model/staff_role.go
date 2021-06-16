package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/datatypes"
)

// StaffRole ...
type StaffRole struct {
	ID           string
	Name         datatypes.JSON
	Role         string         `gorm:"primaryKey"`
	Permissions  pq.StringArray `gorm:"type:text[]"`
	SearchTokens TsVector
	CreatedAt    time.Time
}

// TableName overrides the table name
func (StaffRole) TableName() string {
	return "staff_roles"
}
