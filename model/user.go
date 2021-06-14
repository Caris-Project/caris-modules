package model

import (
	"time"

	"gorm.io/datatypes"
)

// User ...
type User struct {
	ID                     string `gorm:"primaryKey"`
	IdentityCode           string `gorm:"uniqueIndex"`
	Name                   string
	PhoneCountryCode       string
	PhoneNumber            string
	PhoneFull              string
	IsPhoneVerified        bool
	PhoneVerifiedAt        time.Time
	Email                  string
	Avatar                 string
	Status                 string
	RegisterFrom           string
	Facebook               datatypes.JSON // UserFacebook
	Google                 datatypes.JSON
	Apple                  datatypes.JSON
	SearchTokens           TsVector
	IsNewUser              bool
	LastViewNotificationAt time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time

	// Ref
	Province   LocationProvince `gorm:"foreignKey:ProvinceID"`
	ProvinceID int
}

// TableName overrides the table name
func (User) TableName() string {
	return "users"
}

// UserFacebook ...
type UserFacebook struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserStats ...
type UserStats struct {
	ID              string `gorm:"primaryKey"`
	Post            int
	CurrentCoin     int
	Reward          int
	Car             int
	Follower        int
	Following       int
	ReferralTotal   int
	ReferralSuccess int

	// Ref
	User   User `gorm:"foreignKey:UserID"`
	UserID string
}

// TableName overrides the table name
func (UserStats) TableName() string {
	return "user_stats"
}
