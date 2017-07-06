// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// lib
package sql

import "database/sql"

// NullString represents a string that may be null.
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid: true,
	}
}

func CheckNullString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

// NullInt64 represents an int64 that may be null.
func NewNullInt64(s int64) sql.NullInt64 {
	if s == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: s,
		Valid: true,
	}
}