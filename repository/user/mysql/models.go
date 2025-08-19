package mysqluserrepo

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type mySQLUser struct {
	iD        int
	eMail     string
	totalTime int
}

type mySQLUserRepository struct {
	db *sql.DB

	mu   sync.Mutex
	stmt map[string]*sql.Stmt
}
