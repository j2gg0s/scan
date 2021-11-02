package scan

import (
	"reflect"

	"github.com/j2gg0s/scan/schema"
)

type mapModel struct {
	dialect schema.Dialect
	dest    *map[string]interface{}
}

func (m *mapModel) Fields(cols []Column) ([]interface{}, error) {
	fds := make([]interface{}, len(cols))

	dest := *m.dest
	for i, col := range cols {
		val := reflect.New(col.Type)
		dest[col.Name] = val.Elem()
		fds[i] = val.Interface()
	}

	return fds, nil
}

func newMapModel(dialect schema.Dialect, dest *map[string]interface{}) *mapModel {
	return &mapModel{
		dialect: dialect,
		dest:    dest,
	}
}
