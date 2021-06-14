package audit

import (
	"errors"
	"fmt"

	"github.com/Caris-Project/caris-modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBConfig ...
type MongoDBConfig struct {
	Host, User, Password, DBName, Mechanism, Source string
}

// Config ...
type Config struct {
	// Source of server, e.g: caris
	Source string
	// Target of log
	Targets []string
	// MongoDB config, for save documents
	MongoDB MongoDBConfig
}

// Service ...
type Service struct {
	Config
	DB *mongo.Database
}

// NewInstance ...
func NewInstance(config Config) (*Service, error) {
	if config.Source == "" || config.MongoDB.Host == "" {
		return nil, errors.New("please provide all information that needed: source, mongodb")
	}

	// Connect MongoDB
	err := mongodb.Connect(
		config.MongoDB.Host,
		config.MongoDB.User,
		config.MongoDB.Password,
		config.MongoDB.DBName,
		config.MongoDB.Mechanism,
		config.MongoDB.Source,
	)
	if err != nil {
		fmt.Println("Cannot init module AUDIT", err)
		return nil, err
	}

	s := Service{
		Config: config,
		DB:     mongodb.GetInstance(),
	}

	// index mongo
	s.indexDB()

	return &s, nil
}

// getColName ...
func getColName(source, target string) string {
	return source + "-" + target
}
