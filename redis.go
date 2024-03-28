package testx

import (
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
)

// TestRedis is the interface for database operations.
type TestRedis struct {
	client redis.Cmdable
}

// NewTestRedis creates a new TestRedis instance.
func NewTestRedis(env string, t *testing.T) *TestRedis {
	dsn := os.Getenv(env)
	if dsn == "" {
		t.Skipf("skip redis test, env is: %s", env)
	}
	options, err := redis.ParseURL(dsn)
	if err != nil {
		return nil
	}
	client := redis.NewClient(options)
	return &TestRedis{client}
}
