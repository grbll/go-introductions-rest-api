package repository

import (
	"context"
	"database/sql"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/grbll/go-introductions-rest-api/models"
)

var getUserByIdSql string = "SELECT * FROM users WHERE user_id = ? LIMIT 1"

type MySQLUserRepository struct {
	db *sql.DB

	mu sync.Mutex

	userByIdStmt *sql.Stmt
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db, mu: sync.Mutex{}}
}

func (r *MySQLUserRepository) ensureUserByIdStatement(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.userByIdStmt != nil {
		return nil
	}

	stmt, err := r.db.PrepareContext(ctx, getUserByIdSql)
	if err != nil {
		return err
	}

	r.userByIdStmt = stmt
	return nil
}

func (r *MySQLUserRepository) GetUserById(ctx context.Context, userid int) (*User, error) {
	ensureCtx, ensureCancel := context.WithTimeout(ctx, time.Second*2)
	defer ensureCancel()

	err := r.ensureUserByIdStatement(ensureCtx)
	if err != nil {
		return nil, err
	}

	var user *User = &User{}

	queryCtx, queryCancel := context.WithTimeout(ctx, time.Second*2)
	defer queryCancel()

	err = r.userByIdStmt.QueryRowContext(queryCtx, userid).Scan(&user.UserId, &user.Email, &user.TotalTime)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *MySQLUserRepository) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.userByIdStmt != nil {
		return r.userByIdStmt.Close()
	}

	return nil
}
