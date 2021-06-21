package queue

// Init ...
func Init(redisURI, redisPassword string) {
	newClient(redisURI, redisPassword)
	newWorker(redisURI, redisPassword)
}
