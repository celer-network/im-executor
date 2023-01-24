package dal

import (
	"fmt"
	"strings"
)

const (
	AND = "AND"
	OR  = "OR"
)

type WhereBuilder struct {
	exprs []string
}

func (w *WhereBuilder) And(expr string, args ...interface{}) *WhereBuilder {
	if len(w.exprs) > 0 {
		w.add(AND)
	}
	return w.add(expr, args...)
}

func (w *WhereBuilder) Or(expr string, args ...interface{}) *WhereBuilder {
	if len(w.exprs) > 0 {
		w.add(OR)
	}
	return w.add(expr, args...)
}

func (w *WhereBuilder) Build() string {
	return w.String()
}

func (w *WhereBuilder) String() string {
	return strings.Join(w.exprs, " ")
}

func (w *WhereBuilder) add(expr string, args ...interface{}) *WhereBuilder {
	w.exprs = append(w.exprs, fmt.Sprintf(expr, args...))
	return w
}
