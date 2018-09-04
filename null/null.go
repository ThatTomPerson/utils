package null

import (
	"database/sql"
)

// String returns a sql.NullString
func String(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// Int returns a sql.NullInt64
func Int(s interface{}) sql.NullInt64 {
	var i int64
	switch a := s.(type) {
	case int:
		i = int64(a)
	case int8:
		i = int64(a)
	case int16:
		i = int64(a)
	case int32:
		i = int64(a)
	case int64:
		i = a
	case uint:
		i = int64(a)
	case uint8:
		i = int64(a)
	case uint16:
		i = int64(a)
	case uint32:
		i = int64(a)
	case uint64:
		i = int64(a)
	default:
		return sql.NullInt64{Valid: false}
	}

	return sql.NullInt64{Int64: i, Valid: i != 0}
}

// Bool returns a sql.NullBool
func Bool(b interface{}) sql.NullBool {
	if a, ok := b.(bool); ok {
		return sql.NullBool{
			Bool:  a,
			Valid: true,
		}
	}

	return sql.NullBool{
		Valid: false,
	}
}

// Float returns a sql.NullFloat64
func Float(f interface{}) sql.NullFloat64 {
	var i float64
	switch a := f.(type) {
	case float32:
		i = float64(a)
	case float64:
		i = a
	default:
		return sql.NullFloat64{Valid: false}
	}

	return sql.NullFloat64{Float64: i, Valid: i != 0.0}
}
