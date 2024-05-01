package repository

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/aicacia/auth/api/app/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func InitDB() error {
	connection, err := sqlx.Connect("postgres", env.GetDatabaseUrl())
	if err != nil {
		return err
	}
	db = connection
	return nil
}

func CloseDB() error {
	if db == nil {
		return nil
	}
	err := db.Close()
	if err != nil {
		return err
	}
	db = nil
	return nil
}

func NullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{Valid: true, String: *s}
}

func StringFromSQLNullString(null sql.NullString) *string {
	if null.Valid {
		return &null.String
	}
	return nil
}

func Int32FromSQLInt32String(null sql.NullInt32) *int32 {
	if null.Valid {
		return &null.Int32
	}
	return nil
}

func TimeFromSQLNullTime(nullTime sql.NullTime) *time.Time {
	if nullTime.Valid {
		return &nullTime.Time
	}
	return nil
}

func ValidConnection() bool {
	return db.Ping() == nil
}

func All[T any](query string, args ...interface{}) ([]T, error) {
	var rows []T
	err := db.Select(&rows, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func NamedAll[T any](query string, named interface{}) ([]T, error) {
	rows, err := db.NamedQuery(query, named)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []T
	for rows.Next() {
		var row T
		err := rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func GetOptional[T any](query string, args ...interface{}) (*T, error) {
	rows, err := All[T](query, args...)
	if err != nil {
		return nil, err
	}
	if len(rows) > 0 {
		return &rows[0], nil
	}
	return nil, nil
}

func NamedGetOptional[T any](query string, named interface{}) (*T, error) {
	rows, err := db.NamedQuery(query, named)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		var row T
		err := rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		return &row, nil
	}
	return nil, nil
}

func Get[T any](query string, args ...interface{}) (T, error) {
	var row T
	err := db.Get(&row, query, args...)
	if err != nil {
		return row, err
	}
	return row, nil
}

func NamedGet[T any](query string, named interface{}) (T, error) {
	var row T
	rows, err := db.NamedQuery(query, named)
	if err != nil {
		return row, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.StructScan(&row)
		if err != nil {
			return row, err
		}
		return row, nil
	}
	return row, sql.ErrNoRows
}

func Execute(query string, args ...interface{}) (bool, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

func ExecuteNamed(query string, named interface{}) (bool, error) {
	result, err := db.NamedExec(query, named)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

func Transaction[T any](fn func(tx *sqlx.Tx) (T, error)) (T, error) {
	tx, txErr := db.Beginx()
	if txErr != nil {
		return *new(T), txErr
	}
	result, fnErr := fn(tx)
	if fnErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return *new(T), errors.Join(fnErr, rollbackErr)
		}
		return *new(T), fnErr
	} else {
		commitErr := tx.Commit()
		if commitErr != nil {
			return *new(T), commitErr
		}
		return result, nil
	}
}

func IsDuplicateKeyError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint")
}
