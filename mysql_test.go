package testx

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mysqlTestSuite struct {
	*suite.Suite
	TestDB
}

func TestMySQLSuite(t *testing.T) {
	suite.Run(t, &mysqlTestSuite{
		Suite:  new(suite.Suite),
		TestDB: NewTestMySQL("MYSQL_DSN", t),
	})
}

func (s *mysqlTestSuite) Test_Create() {
	// test MySQL operations
}
