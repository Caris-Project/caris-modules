package model

import (
	"strings"
	"time"

	"github.com/Caris-Project/caris-modules/utils"
)

// Staff ...
type Staff struct {
	ID             string `gorm:"primaryKey"`
	Phone          string `gorm:"uniqueKey"`
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

// GenerateSearchTokens ...
func (s *Staff) GenerateSearchTokens() {
	values := []string{utils.RemoveDiacritics(s.Name), s.Phone}
	vecValue := strings.Join(values, " ")
	s.SearchTokens = TsVector{Value: vecValue}
}
