package driver_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/tomarrell/lbadd/driver"
)

var (
	LocalDatabaseAddress = ":53672"
)

func TestMain(m *testing.M) {
	// TODO(TimSatke): start local database or mock it

	exitCode := m.Run()

	// TODO(TimSatke): shut down local database or destroy mock

	os.Exit(exitCode)
}

func TestDriverRegister(t *testing.T) {
	t.Logf("local database address: %v", LocalDatabaseAddress)
	assert.Contains(t, sql.Drivers(), "lbadd")
}

func TestStatement(t *testing.T) {
	t.SkipNow() // skip until database is functional

	assert := assert.New(t)

	pool, err := sql.Open("lbadd", LocalDatabaseAddress)
	assert.NoError(err)
	assert.NoError(pool.Ping())
	defer func() {
		assert.NoError(pool.Close())
	}()

	stmt, err := pool.Prepare(`CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(25));`)
	assert.NoError(err)

	_, err = stmt.Exec()
	assert.NoError(err)
	assert.NoError(stmt.Close())

	stmt, err = pool.Prepare(`INSERT INTO users (name) VALUES (?);`)
	assert.NoError(err)

	result, err := stmt.Exec("jdoe")
	assert.NoError(err)

	affected, err := result.RowsAffected()
	assert.NoError(err)
	assert.Equal(affected, 1)

	assert.NoError(stmt.Close())
}
