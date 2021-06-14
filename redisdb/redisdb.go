package redisdb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/logrusorgru/aurora"

	"github.com/Caris-Project/caris-modules/logger"
)

var (
	rdb *redis.Client
)

// Connect ...
func Connect(uri, password string) error {
	ctx := context.Background()

	rdb = redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: password,
		DB:       0, // use default DB
	})

	// Test
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error when connect to Redis", uri, err)
		return err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO REDIS: " + uri))

	return nil
}

// SetKeyValue ...
func SetKeyValue(ctx context.Context, key string, value interface{}) {
	storeByte, _ := json.Marshal(value)
	rdb.Set(ctx, key, storeByte, 0)
}

// GetValueByKey ...
func GetValueByKey(ctx context.Context, key string) string {
	value, _ := rdb.Get(ctx, key).Result()
	return value
}

// DelKey ...
func DelKey(ctx context.Context, key string) {
	rdb.Del(ctx, key)
}

// DelWithPrefix ...
func DelWithPrefix(ctx context.Context, prefix string) {
	iter := rdb.Scan(ctx, 0, prefix, 0).Iterator()
	for iter.Next(ctx) {
		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			logger.Error("Cannot delete redis key with prefix", logger.LogData{
				"prefix": prefix,
			})
		}
	}
	if err := iter.Err(); err != nil {
		logger.Error("Cannot (LOOP) delete redis key with prefix", logger.LogData{
			"prefix": prefix,
		})
	}
}
