package model

import "time"

// Referral ...
type Referral struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	InviteeID string
	Success   bool
	CreatedAt time.Time
	UpdatedAt time.Time

	// Ref
	User    User `gorm:"foreignKey:UserID"`
	Invitee User `gorm:"foreignKey:InviteeID"`
}

// TableName overrides the table name
func (Referral) TableName() string {
	return "user_referrals"
}
