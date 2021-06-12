package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect ...
func Connect(host, user, password, dbname, port string) (*gorm.DB, error) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC() // Set db timezone to UTC
		},
	})

	if err != nil {
		fmt.Println("Error when connect to PostgreSQL database", dsn, err)
		return nil, err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO POSTGRESQL: " + dsn))

	// Config connection pool
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}

// GetInstance ...
func GetInstance(ctx context.Context) *gorm.DB {
	return db.Session(&gorm.Session{
		SkipHooks: true,
		Context:   ctx,
	})
}
