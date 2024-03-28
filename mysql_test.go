package testx

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
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
	ctx := context.Background()
	_, err := s.Get().ExecContext(ctx, "create table if not exists `test_table` (name varchar(255))")
	s.Require().NoError(err)

	type counter struct {
		C int `db:"c"`
	}
	type testTable struct {
		Name string `db:"name"`
	}

	s.MustTransaction(ctx, s.Suite, func(ctx context.Context, tx *sqlx.Tx) {
		_, err = tx.ExecContext(ctx, "INSERT INTO test_table (name) VALUES (?)", "test")
		s.Require().NoError(err)

		var count counter
		err = tx.GetContext(ctx, &count, "SELECT count(*) as c FROM `test_table`")
		s.Require().NoError(err)
		s.Require().Equal(1, count.C)

		var tt testTable
		err = tx.GetContext(ctx, &tt, "SELECT * FROM `test_table`")
		s.Require().NoError(err)
		s.Require().Equal("test", tt.Name)
	})
}
