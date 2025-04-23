package redis

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitRedis() error {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatal("invalid REDIS_DB value")
		return errors.Wrap(err, "invalid REDIS_DB value")
	}

	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"), // kosong jika tidak ada password
		DB:       db,
	})

	// Test connection
	_, err = RDB.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("failed to connect to Redis")
		return errors.Wrap(err, "failed to connect to Redis")
	}

	return nil
}

func SetToRedisWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
    return RDB.Set(ctx, key, value, ttl).Err()
}

// GetFromRedis mengambil data dari Redis berdasarkan key
func GetFromRedis(ctx context.Context, key string) (string, error) {
    result, err := RDB.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", nil // Key tidak ditemukan
    }
    if err != nil {
        return "", err
    }
    return result, nil
}

// DeleteFromRedis menghapus data dari Redis berdasarkan key
func DeleteFromRedis(ctx context.Context, key string) error {
    return RDB.Del(ctx, key).Err()
}
