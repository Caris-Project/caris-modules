package queue

import (
	"log"

	"github.com/hibiken/asynq"
)

var mux *asynq.ServeMux

// newWorker ...
func newWorker(redisURI, redisPassword string) {
	// Init redis connection
	redisConn := asynq.RedisClientOpt{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       0,
	}

	// Init worker
	worker := asynq.NewServer(redisConn, asynq.Config{
		Concurrency: 100,
		Queues: map[string]int{
			PriorityCritical: 6,
			PriorityDefault:  3,
			PriorityLow:      1,
		},
	})

	// Init mux server
	mux = asynq.NewServeMux()

	// Run server
	go func() {
		if err := worker.Run(mux); err != nil {
			log.Fatalln(err)
		}
	}()
}

// GetMux ...
func GetMux() *asynq.ServeMux {
	return mux
}
