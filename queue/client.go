package queue

import (
	"github.com/hibiken/asynq"
)

var client *asynq.Client

// newClient ...
func newClient(redisURI, redisPassword string) {
	// Init redis connection
	redisConn := asynq.RedisClientOpt{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       0,
	}

	// Init client
	client = asynq.NewClient(redisConn)
}

// GetClient ...
func GetClient() *asynq.Client {
	return client
}
