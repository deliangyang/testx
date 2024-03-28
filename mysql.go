package testx

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"

	// load MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// TestDB is the interface for database operations.
type TestDB interface {
	Get() *sqlx.DB
	MustTransaction(ctx context.Context, s *suite.Suite, fn func(ctx context.Context, tx *sqlx.Tx))
}

var _ TestDB = &TestMySQL{}

// TestMySQL wraps the database connection.
type TestMySQL struct {
	db *sqlx.DB
}

// Get returns the database connection.
func (t *TestMySQL) Get() *sqlx.DB {
	return t.db
}

// NewTestMySQL creates a new TestDB instance.
func NewTestMySQL(env string, t *testing.T) *TestMySQL {
	dsn := os.Getenv(env)
	if dsn == "" {
		t.Skipf("skip db test, env is: %s", env)
	}
	db := sqlx.MustConnect("mysql", dsn)

	return &TestMySQL{db: db}
}

// MustTransaction wraps the transaction rollback logic.
func (t *TestMySQL) MustTransaction(ctx context.Context, s *suite.Suite, fn func(ctx context.Context, tx *sqlx.Tx)) {
	tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{})
	s.Require().NoError(err)
	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()
		s.Require().NoError(err)
	}(tx)
	fn(ctx, tx)
}
