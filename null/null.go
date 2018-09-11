// Package null reduces some boilerplate when working with sql.Null types
package null

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func Time(s interface{}) mysql.NullTime {
	var v mysql.NullTime
	v.Scan(s)
	return v
}

// String returns a sql.NullString
func String(s interface{}) sql.NullString {
	var v sql.NullString
	v.Scan(s)
	return v
}

// Int returns a sql.NullInt64
func Int(s interface{}) sql.NullInt64 {
	var v sql.NullInt64
	v.Scan(s)
	return v
}

// Float returns a sql.NullFloat64
func Float(s interface{}) sql.NullFloat64 {
	var v sql.NullFloat64
	v.Scan(s)
	return v
}

// Bool returns a sql.NullBool
func Bool(s interface{}) sql.NullBool {
	var v sql.NullBool
	v.Scan(s)
	return v
}
