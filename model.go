package scan

import (
	"fmt"
	"reflect"

	"github.com/j2gg0s/scan/schema"
)

var (
	DiscardUnknownColumn bool = true
)

func newModel(dialect schema.Dialect, dest interface{}) (Model, error) {
	v := reflect.ValueOf(dest)

	if !v.IsValid() {
		return nil, fmt.Errorf("got nil model")
	}
	if v.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("want %s, got %s", reflect.Ptr, v.Kind())
	}

	v = v.Elem()
	switch v.Kind() {
	case reflect.Map:
		return newMapModel(dialect, dest.(*map[string]interface{})), nil
	case reflect.Struct:
		return newStructModel(dialect, dest, v), nil
	case reflect.Slice:
		switch elemType := sliceElemType(v); elemType.Kind() {
		case reflect.Struct:
			return newSliceModel(dialect, dest, v, elemType), nil
		case reflect.Map:
			return newMapSliceModel(dialect, dest.(*[]map[string]interface{})), nil
		}
	}
	return nil, nil
}

type Model interface {
	Fields([]Column) ([]interface{}, error)
}

type Column struct {
	Name string
	Type reflect.Type
}

type Rows interface {
	Columns() ([]Column, error)
	Next() bool
	Scan(...interface{}) error
	Err() error
}

func Scan(dialect schema.Dialect, rows Rows, dest interface{}) error {
	model, err := newModel(dialect, dest)
	if err != nil {
		return err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil
	}

	for rows.Next() {
		fds, err := model.Fields(cols)
		if err != nil {
			return err
		}
		if err := rows.Scan(fds...); err != nil {
			return err
		}
	}

	return rows.Err()
}
