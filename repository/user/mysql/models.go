package mysqluserrepo

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLUser struct {
	ID        int
	EMail     string
	TotalTime int
}

type MySQLUserRepository struct {
	db *sql.DB

	mu   sync.Mutex
	stmt map[string]*sql.Stmt
}
