package mysqluserrepo

import (
	"context"
	"errors"
	"sync"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLUserRepository(db *sql.DB) *mySQLUserRepository {
	return &mySQLUserRepository{db: db, mu: sync.Mutex{}, stmt: map[string]*sql.Stmt{}}
}

func (r *mySQLUserRepository) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var errorCollection []error

	for name, stmt := range r.stmt {
		if stmt != nil {
			if err := stmt.Close(); err != nil {
				errorCollection = append(errorCollection, err)
			}
			delete(r.stmt, name)
		}
	}

	if len(errorCollection) > 0 {
		return errors.Join(errorCollection...)
	}

	return nil
}

// func (r *MySQLUserRepository) GetById(ctx context.Context, userid int) (*User, error) {
// 	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
// 	defer cancel()
//
// 	stmt, err := r.getStmt(ctx, getById)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	rows, err := stmt.QueryContext(ctx, userid)
// 	defer rows.Close()
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	user, err := newUserFromRow(rows)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return user, nil
// }

func (r *mySQLUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	stmt, err := r.getStmt(ctx, existsByEmail)
	if err != nil {
		return false, err
	}

	var exists bool

	err = stmt.QueryRowContext(ctx, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *mySQLUserRepository) InsertUser(ctx context.Context, email string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	stmt, err := r.getStmt(ctx, insertUser)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
