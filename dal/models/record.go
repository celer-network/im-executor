package models

import "strings"

type Record interface {
	Columns() Columns
	Scan(row Scanner) error
}

type Columns []string

func (c Columns) String() string {
	return strings.Join(c, ",")
}

type Scanner interface {
	Scan(dest ...interface{}) error
}
