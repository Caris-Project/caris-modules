package model

import (
	"time"

	"gorm.io/datatypes"
)

// Notification ...
type Notification struct {
	ID        string
	Category  string
	Type      string
	TargetID  string
	Data      datatypes.JSON // Custom format for each type of notification
	IsRead    bool
	CreatedAt time.Time

	// Ref
	User   User `gorm:"foreignKey:UserID"`
	UserID string
}

// TableName overrides the table name
func (Notification) TableName() string {
	return "notifications"
}

// NotificationData ...
type NotificationData struct {
	Status   string `json:"status,omitempty"`
	UserID   string `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
	Photo    string `json:"photo,omitempty"`
}

// NotificationAction ...
type NotificationAction struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// NotificationPushData ...
type NotificationPushData struct {
	ID      string              `json:"id"`
	Title   string              `json:"title"`
	Content string              `json:"content"`
	Action  *NotificationAction `json:"action"`
	Photo   string              `json:"photo"`
}

// NotificationResponse ...
type NotificationResponse struct {
	ID      string              `json:"id"`
	Title   string              `json:"title"`
	Content string              `json:"content"`
	Action  *NotificationAction `json:"action"`
	Photo   string              `json:"photo"`
	Data    *NotificationData   `json:"-"`
}
