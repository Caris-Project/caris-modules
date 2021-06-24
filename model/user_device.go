package model

import "time"

// UserDevice ...
type UserDevice struct {
	ID              string `gorm:"primaryKey"`
	DeviceID        string `gorm:"uniqueIndex"`
	FCMToken        string `gorm:"uniqueIndex"`
	OSName          string
	OSVersion       string
	AppVersion      string
	IP              string
	IsMobile        bool
	Language        string
	CreatedAt       time.Time
	LastActivatedAt time.Time

	// Ref
	User   User `gorm:"UserID"`
	UserID string
}

// TableName overrides the table name
func (UserDevice) TableName() string {
	return "user_devices"
}
