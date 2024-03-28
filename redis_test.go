package testx

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testRedisSuite struct {
	*suite.Suite
	*TestRedis
}

// TestRedisSuite runs the Redis test suite.
func TestRedisSuite(t *testing.T) {
	suite.Run(t, &testRedisSuite{
		Suite:     new(suite.Suite),
		TestRedis: NewTestRedis("REDIS_DSN", t),
	})
}
