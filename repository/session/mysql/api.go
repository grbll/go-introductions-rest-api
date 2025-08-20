package mysqlsessionrepo

import (
	"errors"
	"sync"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLSessionRepository(db *sql.DB) *mySQLSessionRepository {
	return &mySQLSessionRepository{db: db, mu: sync.Mutex{}, stmt: map[string]*sql.Stmt{}}
}

func (r *mySQLSessionRepository) Close() error {
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
