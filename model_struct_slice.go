package scan

import (
	"reflect"

	"github.com/j2gg0s/scan/schema"
)

func newStructSliceModel(dialect schema.Dialect, slice interface{}, val reflect.Value, elemType reflect.Type) *structSliceModel {
	return &structSliceModel{
		structModel: structModel{
			dialect: dialect,
			table:   dialect.Tables().Get(elemType),
		},

		slice:    val,
		nextElem: makeSliceNextElemFunc(val),
	}
}

type structSliceModel struct {
	structModel

	slice    reflect.Value
	nextElem func() reflect.Value
}

func (m *structSliceModel) Fields(cols []Column) ([]interface{}, error) {
	return m.fields(m.nextElem(), cols)
}
