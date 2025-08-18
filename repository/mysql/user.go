package mysqlrepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/grbll/go-introductions-rest-api/models"
)

const getById string = "GetById"
const existsByEmail string = "ExistsByEmail"
const insertUser string = "InsertUser"

var queryCollection map[string]string = map[string]string{
	getById:       "SELECT * FROM users WHERE user_id = ? LIMIT 1",
	existsByEmail: "SELECT EXISTS(SELECT 1 FROM users WHERE user_email = ? LIMIT 1)",
	insertUser:    "INSERT INTO users (user_email) VALUES (?)",
}

type MySQLUserRepository struct {
	db *sql.DB

	mu   sync.Mutex
	stmt map[string]*sql.Stmt
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db, mu: sync.Mutex{}, stmt: map[string]*sql.Stmt{}}
}

func (r *MySQLUserRepository) Close() error {
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

func newMappedUser() (*User, map[string]any) {
	var user *User = &User{}
	var mapping map[string]any = map[string]any{}

	mapping["user_id"] = &user.UserId
	mapping["user_email"] = &user.Email
	mapping["user_total_time"] = &user.TotalTime

	return user, mapping
}

func newUserFromRow(rows *sql.Rows) (*User, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var dest []any = make([]any, len(cols), len(cols))
	user, mapping := newMappedUser()
	for i, name := range cols {
		if ref, ok := mapping[name]; ok {
			dest[i] = ref
		}
	}

	err = rows.Scan(dest...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *MySQLUserRepository) getStmt(ctx context.Context, name string) (*sql.Stmt, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if stmt, ok := r.stmt[name]; ok && stmt != nil {
		return stmt, nil
	}

	query, ok := queryCollection[name]
	if !ok {
		return nil, fmt.Errorf("prepare %q: unknown query!", name)
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("prepare %q: %w", name, err)
	}

	r.stmt[name] = stmt
	return stmt, nil
}

func (r *MySQLUserRepository) GetById(ctx context.Context, userid int) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	stmt, err := r.getStmt(ctx, getById)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, userid)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	user, err := newUserFromRow(rows)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *MySQLUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
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

func (r *MySQLUserRepository) InsertUser(ctx context.Context, email string) error {
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
