package schema

import (
	"reflect"

	"github.com/j2gg0s/scan/internal"
	"github.com/j2gg0s/scan/internal/tagparser"
)

type Field struct {
	StructField reflect.StructField

	Type  reflect.Type
	Index []int
	Tag   tagparser.Tag

	Name string

	Scan ScannerFunc
}

func (t *Table) newField(fd reflect.StructField, index []int) *Field {
	tag := tagparser.Parse(fd.Tag.Get(t.dialect.Tag()))

	index = append(index, fd.Index...)
	name := internal.Underscore(fd.Name)
	if tag.Name != "" {
		name = tag.Name
	}

	field := &Field{
		StructField: fd,
		Type:        indirectType(fd.Type),
		Index:       index,
		Tag:         tag,

		Name: name,
	}
	field.Scan = FieldScanner(t.dialect, field)
	return field
}
