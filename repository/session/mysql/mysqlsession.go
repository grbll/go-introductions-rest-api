package mysqlrepository

import (
	"database/sql"
	"errors"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	. "github.com/grbll/go-introductions-rest-api/models"
)

type MySQLSessionRepository struct {
	db *sql.DB

	mu   sync.Mutex
	stmt map[string]*sql.Stmt
}

func NewMySQLSessionRepository(db *sql.DB) *MySQLSessionRepository {
	return &MySQLSessionRepository{db: db, mu: sync.Mutex{}, stmt: map[string]*sql.Stmt{}}
}

func (r *MySQLSessionRepository) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var errorCollection []error = []error{}

	for name, stmt := range r.stmt {
		if stmt != nil {
			if err := stmt.Close(); err != nil {
				errorCollection = append(errorCollection, err)
			}
		}
		delete(r.stmt, name)
	}

	if len(errorCollection) > 0 {
		return errors.Join(errorCollection...)
	}

	return nil
}

func newMappedActiveSession() *ActiveSession {
	return nil
}
