package schema

import (
	"reflect"
)

type Table struct {
	dialect Dialect

	Type     reflect.Type
	Fields   []*Field
	FieldMap map[string]*Field
}

func NewTable(dialect Dialect, typ reflect.Type) *Table {
	t := new(Table)
	t.dialect = dialect
	t.Type = typ

	t.initFields()
	return t
}

func (t *Table) initFields() {
	t.Fields = make([]*Field, 0, t.Type.NumField())
	t.FieldMap = make(map[string]*Field, t.Type.NumField())
	t.addFields(t.Type, nil)
}

func (t *Table) addFields(typ reflect.Type, baseIndex []int) {
	for i := 0; i < typ.NumField(); i++ {
		fd := typ.Field(i)

		unexported := fd.PkgPath != ""

		if unexported && !fd.Anonymous {
			continue
		}

		index := make([]int, len(baseIndex))
		copy(index, baseIndex)

		if fd.Anonymous {
			fieldType := indirectType(fd.Type)
			if fieldType.Kind() != reflect.Struct {
				continue
			}
			t.addFields(fieldType, append(index, fd.Index...))
			continue
		}

		if field := t.newField(fd, index); field != nil {
			t.addField(field)
		}
	}
}

func (t *Table) addField(fd *Field) {
	t.Fields = append(t.Fields, fd)
	t.FieldMap[fd.Name] = fd
}
