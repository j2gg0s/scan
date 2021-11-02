package scan

import "github.com/j2gg0s/scan/schema"

type mapSliceModel struct {
	mapModel

	slice *[]map[string]interface{}
}

func newMapSliceModel(dialect schema.Dialect, slice *[]map[string]interface{}) *mapSliceModel {
	return &mapSliceModel{
		mapModel: mapModel{
			dialect: dialect,
		},
		slice: slice,
	}
}

func (m *mapSliceModel) Fields(cols []Column) ([]interface{}, error) {
	dest := make(map[string]interface{})
	slice := append(*m.slice, dest)
	*m.slice = slice
	m.dest = &dest
	return m.mapModel.Fields(cols)
}
