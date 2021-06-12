package fcm

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
)

// List actions
const (
	LogActionSendByTokens      = "send_by_tokens"
	LogActionSendByTopic       = "send_by_topic"
	LogActionSubscribeTokens   = "subscribe_tokens"
	LogActionUnsubscribeTokens = "unsubscribe_tokens"
)

// Log ...
type Log struct {
	ID           string `gorm:"primaryKey"`
	Action       string
	BatchID      string
	Topics       datatypes.JSON
	TokenCount   int
	SuccessCount int
	FailureCount int
	CreatedAt    time.Time

	tableName string
}

// TableName ...
func (l Log) TableName() string {
	if l.tableName != "" {
		return l.tableName
	}
	return "log_fcm"
}

// Save log to db
func (s Service) saveLog(doc Log) {
	if err := s.DB.Model(Log{}).Create(doc).Error; err != nil {
		fmt.Println("*** Error when create log", err)
		fmt.Println("*** Log", doc)
	}
}
