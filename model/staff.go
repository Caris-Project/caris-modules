package model

import "time"

// Staff ...
type Staff struct {
	ID             string `gorm:"primaryKey"`
	Username       string
	Name           string
	HashedPassword string
	Status         string
	SearchTokens   TsVector
	CreatedAt      time.Time

	// Ref
	Role   StaffRole `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE"`
	RoleID string
}

// TableName overrides the table name
func (Staff) TableName() string {
	return "staffs"
}
