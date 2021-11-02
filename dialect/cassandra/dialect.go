package cassandra

import (
	"github.com/gocql/gocql"

	"github.com/j2gg0s/scan"
	"github.com/j2gg0s/scan/schema"
)

type Dialect struct {
	tables *schema.Tables
}

func (d *Dialect) Name() string {
	return "cassandra"
}

func (d *Dialect) Tag() string {
	return "cassandra"
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

func Scan(iter *gocql.Iter, dest interface{}) error {
	return scan.Scan(Default, &wrapper{iter: iter, scanner: iter.Scanner()}, dest)
}

type wrapper struct {
	iter    *gocql.Iter
	scanner gocql.Scanner
}

func (w *wrapper) Next() bool                     { return w.scanner.Next() }
func (w *wrapper) Scan(dest ...interface{}) error { return w.scanner.Scan(dest...) }
func (w *wrapper) Err() error                     { return w.scanner.Err() }
func (w *wrapper) Columns() ([]scan.Column, error) {
	cqlCols := w.iter.Columns()
	cols := make([]scan.Column, len(cqlCols))
	for i, col := range cqlCols {
		cols[i] = scan.Column{
			Name: col.Name,
			Type: goType(col.TypeInfo),
		}
	}
	return cols, nil
}
