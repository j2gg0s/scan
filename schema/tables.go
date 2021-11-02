package schema

import (
	"fmt"
	"reflect"
	"sync"
)

type Dialect interface {
	Name() string
	Tables() *Tables
	Tag() string
}

type Tables struct {
	dialect Dialect
	tables  sync.Map

	mu sync.Mutex
}

func NewTables(dialect Dialect) *Tables {
	return &Tables{
		dialect: dialect,
	}
}

func (t *Tables) Get(typ reflect.Type) *Table {
	return t.table(typ)
}

func (t *Tables) table(typ reflect.Type) *Table {
	typ = indirectType(typ)
	if typ.Kind() != reflect.Struct {
		panic(fmt.Errorf("got %s, wanted %s", typ.Kind(), reflect.Struct))
	}

	if v, ok := t.tables.Load(typ); ok {
		return v.(*Table)
	}

	t.mu.Lock()
	if v, ok := t.tables.Load(typ); ok {
		t.mu.Unlock()
		return v.(*Table)
	}

	table := NewTable(t.dialect, typ)
	table.initFields()
	t.tables.Store(typ, table)
	t.mu.Unlock()

	return table
}
