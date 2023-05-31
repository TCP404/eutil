package packer

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

const (
	timeTime = "time.Time"
)

type Getter interface {
	Get(path string) Agency
}

type Setter interface {
	Set(key string, value any)
}

type Marshaller interface {
	Marshal() []byte
	Unmarshal([]byte) Any
}

type Agency interface {
	ValueType() ValueType
	Val() reflect.Value
	Idx() []int
	Kind() reflect.Kind
	ToString() string
	ToBool() bool
	ToInt() int
	ToInt32() int32
	ToInt64() int64
	ToUint() uint
	ToUint32() uint32
	ToUint64() uint64
	ToFloat32() float32
	ToFloat64() float64
	Set(any) Agency
	LastError() error
}

// Any generic object representation.
type Any interface {
	Getter
	Setter
	Marshaller
	LastError() error
}

type BaseAny struct {
	origin  any
	originV reflect.Value
	td      TypeDescriptor
	err     error
}

func Parse[T any](in *T, tags ...string) Any {
	td, err := DescribeType(in, tags...)
	if err != nil {
		return &BaseAny{err: err}
	}
	return ParseWithTD(in, td)
}

func ParseWithTD[T any](in *T, td TypeDescriptor) Any {
	k := reflect.TypeOf(in)
	if k.Kind() != reflect.Ptr {
		return &BaseAny{err: errors.New("you must pass in a pointer to a struct")}
	}
	if k.Elem().Kind() != reflect.Struct {
		return &BaseAny{err: errors.New("you must pass in a pointer to a struct")}
	}

	return &BaseAny{
		origin:  in,
		originV: reflect.Indirect(reflect.ValueOf(in)),
		td:      td,
	}
}

func (t *BaseAny) Marshal() []byte {
	var b []byte
	b, t.err = json.Marshal(t.origin)
	return b
}

func (t *BaseAny) Unmarshal(b []byte) (result Any) {
	result = t
	t.err = json.Unmarshal(b, t.origin)
	return t
}

func (t *BaseAny) Get(path string) Agency {
	a := &Agent{}
	p, ok := t.td[path]
	if !ok {
		t.err = errors.Errorf("")
		return a
	}
	a.path = path
	a.idx = p.idx
	a.kind = p.kind
	a.val = t.originV.FieldByIndex(p.idx)
	return a
}

func (t *BaseAny) Set(path string, value any) {
	defer func() {
		if err := recover(); err != nil {
			t.err = errors.New(cast.ToString(err))
		}
	}()

	t.Get(path).Set(value)
}

func (t *BaseAny) LastError() error {
	return t.err
}

type Agent struct {
	fieldInfo
	val  reflect.Value
	path string
	err  error
}

func (a *Agent) ValueType() ValueType {
	k := a.kind
	if a.kind == reflect.Ptr {
		k = a.val.Elem().Kind()
	}
	switch k {
	case reflect.String:
		return StringValue

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return NumberValue

	case reflect.Bool:
		return BoolValue

	case reflect.Array, reflect.Slice:
		return ArrayValue

	case reflect.Struct, reflect.Map:
		return ObjectValue

	case reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return NullableValue

	default:
		return InvalidValue
	}
}
func (a *Agent) Val() reflect.Value {
	return a.val
}
func (a *Agent) Idx() []int {
	return a.idx
}
func (a *Agent) Kind() reflect.Kind {
	return a.kind
}

func (a *Agent) Index(i int) (result Agency) {
	result = a
	if a.ValueType() != ArrayValue {
		a.err = errors.Errorf("%v is not a array or slice", a.path)
		return a
	}
	if l := a.val.Len(); i >= l {
		a.err = errors.Errorf("index out of range. length of %v: %v", a.path, l)
		return a
	}
	v := a.val.Index(i)
	e := &Agent{}
	e.path = a.path + "#" + strconv.Itoa(i)
	e.idx = a.idx
	e.kind = v.Kind()
	e.val = v

	return a
}
func (a *Agent) Set(value any) (result Agency) {
	result = a
	defer func() {
		if err := recover(); err != nil {
			a.err = errors.New(cast.ToString(err))
		}
	}()

	v := reflect.ValueOf(value)
	switch {
	case a.kind == reflect.Pointer && v.Type().Kind() != reflect.Pointer:
		{
			a.err = errors.Errorf("you must pass in a pointer, cause type of %v is pointer", a.path)
		}
	case a.kind != reflect.Pointer && v.Type().Kind() == reflect.Pointer:
		{
			a.val.Set(reflect.Indirect(v))
		}
	default:
		{
			a.val.Set(v)
		}
	}
	return a
}

func (a *Agent) ToString() string {
	v := reflect.Indirect(a.val)
	if v.IsZero() {
		return ""
	}
	if v.Type().String() == timeTime {
		return v.Interface().(time.Time).String()
	}
	return v.String()
}
func (a *Agent) ToBool() (b bool) {
	if a.kind == reflect.Bool {
		b = a.val.Bool()
	}
	return
}
func (a *Agent) ToInt() (i int) {
	switch a.kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if a.val.CanInt() {
			i = int(a.val.Int())
		}
	}
	return
}
func (a *Agent) ToInt32() (i int32) {
	if a.val.CanInt() {
		i = int32(a.val.Int())
	}
	return
}
func (a *Agent) ToInt64() (i int64) {
	if a.val.CanInt() {
		i = a.val.Int()
	}
	return
}
func (a *Agent) ToUint() (u uint) {
	if a.val.CanUint() {
		u = uint(a.val.Int())
	}
	return
}
func (a *Agent) ToUint32() (u uint32) {
	if a.val.CanUint() {
		u = uint32(a.val.Int())
	}
	return
}
func (a *Agent) ToUint64() (u uint64) {
	if a.val.CanUint() {
		u = uint64(a.val.Int())
	}
	return
}
func (a *Agent) ToFloat32() (f float32) {
	if a.val.CanFloat() {
		f = float32(a.val.Float())
	}
	return
}
func (a *Agent) ToFloat64() (f float64) {
	if a.val.CanFloat() {
		f = a.val.Float()
	}
	return
}

func (a *Agent) LastError() error {
	return a.err
}
