package util

import "github.com/jackc/pgx/v5/pgconn"

func IsUniqueConstraintError(err error) bool {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		return pgErr.Code == "23505"
	}
	
	return false
}