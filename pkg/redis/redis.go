package redis

import (
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"github.com/anhnmt/golang-gateway-boilerplate/internal/bootstrap/config"
)

const (
	defaultTimeout = 10 * time.Second
)

// NewRedis is new redis.
func NewRedis(ctx context.Context) (redis.UniversalClient, error) {
	if !config.RedisEnabled() {
		return nil, nil
	}

	addrs := strings.Split(config.RedisAddress(), ",")
	db := config.RedisDB()

	log.Info().
		Strs("redis_url", addrs).
		Int("redis_db", db).
		Msg("Connecting to Redis")

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    addrs,
		Password: config.RedisPassword(),
		DB:       db, // use default DB
		PoolSize: 100,
	})

	newCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	if err := client.Ping(newCtx).Err(); err != nil {
		return nil, err
	}

	log.Info().Msg("Connecting to Redis successfully.")
	return client, nil
}
