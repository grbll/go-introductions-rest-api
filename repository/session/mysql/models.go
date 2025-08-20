package mysqlsessionrepo

import (
	"sync"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type mySQLSessionRepository struct {
	db *sql.DB

	mu   sync.Mutex
	stmt map[string]*sql.Stmt
}
