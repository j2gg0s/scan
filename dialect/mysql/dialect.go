package mysql

import (
	"database/sql"
	"reflect"

	"github.com/j2gg0s/scan"
	"github.com/j2gg0s/scan/schema"
)

type Dialect struct {
	tables *schema.Tables
}

func (d *Dialect) Name() string {
	return "mysql"
}

func (d *Dialect) Tag() string {
	return "mysql"
}

func (d *Dialect) Tables() *schema.Tables {
	return d.tables
}

func New() *Dialect {
	d := new(Dialect)
	d.tables = schema.NewTables(d)
	return d
}

var Default = New()

func Scan(rows *sql.Rows, dest ...interface{}) error {
	return scan.Scan(Default, &wrapper{rows: rows}, dest...)
}

type wrapper struct {
	rows *sql.Rows
}

func (w *wrapper) Next() bool                     { return w.rows.Next() }
func (w *wrapper) Scan(dest ...interface{}) error { return w.rows.Scan(dest...) }
func (w *wrapper) Err() error                     { return w.rows.Err() }
func (w *wrapper) Columns() ([]scan.Column, error) {
	names, err := w.rows.Columns()
	if err != nil {
		return nil, err
	}

	types, err := w.rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	cols := make([]scan.Column, len(names))
	for i, name := range names {
		cols[i] = scan.Column{
			Name: name,
		}
		switch types[i].DatabaseTypeName() {
		case "VARCHAR", "CHAR":
			cols[i].Type = reflect.TypeOf("")
		default:
			cols[i].Type = types[i].ScanType()
		}
	}

	return cols, nil
}
