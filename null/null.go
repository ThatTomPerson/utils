package null

import (
	"database/sql"
)

func String(s interface{}) sql.NullString {
	var v sql.NullString
	v.Scan(s)
	return v
}
func Int(s interface{}) sql.NullInt64 {
	var v sql.NullInt64
	v.Scan(s)
	return v
}
func Float(s interface{}) sql.NullFloat64 {
	var v sql.NullFloat64
	v.Scan(s)
	return v
}
func Bool(s interface{}) sql.NullBool {
	var v sql.NullBool
	v.Scan(s)
	return v
}
