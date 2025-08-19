package mysqluserrepo

import (
	"context"
	"fmt"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func newMappedUser() (*MySQLUser, map[string]any) {
	var user *MySQLUser = &MySQLUser{}
	var mapping map[string]any = map[string]any{}

	mapping["user_id"] = &user.ID
	mapping["user_email"] = &user.EMail
	mapping["user_total_time"] = &user.TotalTime

	return user, mapping
}

// func newUserFromRow(rows *sql.Rows) (*MySQLUser, error) {
// 	cols, err := rows.Columns()
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var dest []any = make([]any, len(cols), len(cols))
// 	user, mapping := newMappedUser()
// 	for i, name := range cols {
// 		if ref, ok := mapping[name]; ok {
// 			dest[i] = ref
// 		}
// 	}
//
// 	err = rows.Scan(dest...)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return user, nil
// }

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
