package scan

import (
	"fmt"
	"reflect"

	"github.com/j2gg0s/scan/schema"
)

type sliceModel struct {
	dialect   schema.Dialect
	values    []reflect.Value
	nextElems []func() reflect.Value
}

func newSliceModel(dialect schema.Dialect, values []reflect.Value) *sliceModel {
	m := &sliceModel{
		dialect:   dialect,
		values:    values,
		nextElems: make([]func() reflect.Value, len(values)),
	}
	for i, value := range values {
		m.nextElems[i] = makeSliceNextElemFunc(value)
	}
	return m
}

func (m *sliceModel) Fields(cols []Column) ([]interface{}, error) {
	if len(m.values) != len(cols) {
		return nil, fmt.Errorf("want %d columns, got %d", len(cols), len(m.values))
	}

	fds := make([]interface{}, len(cols))
	for i := 0; i < len(m.values); i++ {
		fds[i] = m.nextElems[i]().Addr().Interface()
	}
	return fds, nil
}
