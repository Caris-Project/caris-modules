package model

import (
	"gorm.io/datatypes"
	"time"
)

// Post ...
type Post struct {
	ID                   string `gorm:"primaryKey"`
	Content              string
	Title                string // system user's post
	Status               string
	FromCommunityExpert  bool
	PushNotification     bool
	PostNow              bool
	UseForNews           bool
	AudioURL             string
	BriefContent         string
	BriefContentAudioURL string
	PublishAt            time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	SearchTokens         TsVector
	Photos               datatypes.JSON
}

// TableName overrides the table name
func (Post) TableName() string {
	return "posts"
}
