package scan

import (
	"fmt"
	"reflect"

	"github.com/j2gg0s/scan/schema"
)

func newStructModel(dialect schema.Dialect, dest interface{}, val reflect.Value) *structModel {
	return &structModel{
		dialect: dialect,
		table:   dialect.Tables().Get(val.Type()),
		dest:    val,
	}
}

type structModel struct {
	dialect schema.Dialect
	table   *schema.Table
	dest    reflect.Value
}

func (m *structModel) Fields(cols []Column) ([]interface{}, error) {
	return m.fields(m.dest, cols)
}

func (m *structModel) fields(dest reflect.Value, cols []Column) ([]interface{}, error) {
	fds := make([]interface{}, len(cols))
	for i, col := range cols {
		fd, ok := m.table.FieldMap[col.Name]
		if ok {
			fds[i] = dest.FieldByIndex(fd.Index).Addr().Interface()
		} else if DiscardUnknownColumn {
			fds[i] = &schema.DiscardColumn{}
		} else {
			return nil, fmt.Errorf("%s doest not have column %s", m.table.Type.Name(), col)
		}
	}

	return fds, nil
}
