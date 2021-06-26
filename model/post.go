package model

import (
	"github.com/Caris-Project/caris-modules/utils"
	"gorm.io/datatypes"
	"strings"
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
	FromSystem           bool
}

// TableName overrides the table name
func (Post) TableName() string {
	return "posts"
}

// GenerateSearchTokens ...
func (p *Post) GenerateSearchTokens() {
	values := []string{utils.RemoveDiacritics(p.Title)}
	vecValue := strings.Join(values, " ")
	p.SearchTokens = TsVector{Value: vecValue}
}

// IsPushNotification ...
func (p *Post) IsPushNotification() bool {
	return p.PushNotification && p.PostNow
}
