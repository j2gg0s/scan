package cassandra

import (
	"math/big"
	"reflect"
	"time"

	"github.com/gocql/gocql"
	"gopkg.in/inf.v0"
)

type UUID [16]byte

type Duration struct {
	Months      int32
	Days        int32
	Nanoseconds int64
}

func goType(t gocql.TypeInfo) reflect.Type {
	switch t.Type() {
	case gocql.TypeVarchar, gocql.TypeAscii, gocql.TypeInet, gocql.TypeText:
		return reflect.TypeOf(*new(string))
	case gocql.TypeBigInt, gocql.TypeCounter:
		return reflect.TypeOf(*new(int64))
	case gocql.TypeTime:
		return reflect.TypeOf(*new(time.Duration))
	case gocql.TypeTimestamp:
		return reflect.TypeOf(*new(time.Time))
	case gocql.TypeBlob:
		return reflect.TypeOf(*new([]byte))
	case gocql.TypeBoolean:
		return reflect.TypeOf(*new(bool))
	case gocql.TypeFloat:
		return reflect.TypeOf(*new(float32))
	case gocql.TypeDouble:
		return reflect.TypeOf(*new(float64))
	case gocql.TypeInt:
		return reflect.TypeOf(*new(int))
	case gocql.TypeSmallInt:
		return reflect.TypeOf(*new(int16))
	case gocql.TypeTinyInt:
		return reflect.TypeOf(*new(int8))
	case gocql.TypeDecimal:
		return reflect.TypeOf(*new(*inf.Dec))
	case gocql.TypeUUID, gocql.TypeTimeUUID:
		return reflect.TypeOf(*new(UUID))
	case gocql.TypeList, gocql.TypeSet:
		return reflect.SliceOf(goType(t.(gocql.CollectionType).Elem))
	case gocql.TypeMap:
		return reflect.MapOf(goType(t.(gocql.CollectionType).Key), goType(t.(gocql.CollectionType).Elem))
	case gocql.TypeVarint:
		return reflect.TypeOf(*new(*big.Int))
	case gocql.TypeTuple:
		// what can we do here? all there is to do is to make a list of interface{}
		tuple := t.(gocql.TupleTypeInfo)
		return reflect.TypeOf(make([]interface{}, len(tuple.Elems)))
	case gocql.TypeUDT:
		return reflect.TypeOf(make(map[string]interface{}))
	case gocql.TypeDate:
		return reflect.TypeOf(*new(time.Time))
	case gocql.TypeDuration:
		return reflect.TypeOf(*new(Duration))
	default:
		return nil
	}
}
