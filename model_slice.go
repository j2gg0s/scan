package scan

import (
	"reflect"

	"github.com/j2gg0s/scan/schema"
)

func newSliceModel(dialect schema.Dialect, slice interface{}, val reflect.Value, elemType reflect.Type) *sliceModel {
	return &sliceModel{
		structModel: structModel{
			dialect: dialect,
			table:   dialect.Tables().Get(elemType),
		},

		slice:    val,
		nextElem: makeSliceNextElemFunc(val),
	}
}

type sliceModel struct {
	structModel

	slice    reflect.Value
	nextElem func() reflect.Value
}

func (m *sliceModel) Fields(cols []Column) ([]interface{}, error) {
	return m.fields(m.nextElem(), cols)
}
