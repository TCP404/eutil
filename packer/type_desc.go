package packer

import (
	"go/ast"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

type fieldPath = string
type fieldInfo = struct {
	idx  []int
	kind reflect.Kind
	val  reflect.StructField
}
type TypeDescriptor map[fieldPath]fieldInfo
func (t TypeDescriptor) GetTagValues(tag string) []string {
	var res []string
	for _, info := range t {
		if v := info.val.Tag.Get(tag); v != "" {
			res = append(res, v)
		}
	}
	return res
}
const DefaultTag = "json"

type structField struct {
	rStructField reflect.Type
	parentPath   []string
	parentIdx    []int
}

func DescribeType[T any](in *T, tags ...string) (TypeDescriptor, error) {
	var (
		queue     = make([]structField, 0)
		descTable = make(map[fieldPath]fieldInfo)
		tag       = DefaultTag
	)
	if len(tags) > 0 {
		tag = tags[0]
	}
	t := reflect.Indirect(reflect.ValueOf(in)).Type()

	// 解析结构体问题核心代码
	// 先将自己放进队列
	queue = append(queue, structField{rStructField: t})
	for len(queue) > 0 {
		// 从队列中弹出一个拿来解析
		var x structField
		x, queue = queue[0], queue[1:]
		x.parentPath = Filter(func(t string) bool { return t != "" }, x.parentPath...)
		err := extract(x.rStructField, x.parentPath, x.parentIdx, tag, &queue, descTable)
		if err != nil {
			return nil, err
		}
	}
	return descTable, nil
}

func extract(t reflect.Type, parentPath []string, parentIdx []int, tag string, queue *[]structField, descTable map[fieldPath]fieldInfo) error {
	if t.Kind() != reflect.Struct {
		return errors.Errorf("you must pass in a pointer to a struct")
	}
	for i, l := 0, t.NumField(); i < l; i++ {
		f := t.Field(i)
		// skip non-exportable naming fields.
		if !ast.IsExported(f.Name) && !f.Anonymous {
			continue
		}
		fieldName := getFieldName(f, tag)

		idx := make([]int, len(parentIdx))
		path := make([]string, len(parentPath))
		copy(idx, parentIdx)
		copy(path, parentPath)
		idx = append(idx, f.Index[0])
		path = append(path, fieldName)

		if f.Type.Kind() != reflect.Struct ||
			f.Type.String() == "time.Time" { // 此处还需处理如 sync.Once 等标准库结构体的问题
			fieldName = strings.Join(path, ".")
			descTable[fieldName] = fieldInfo{idx: idx, kind: f.Type.Kind()}
			continue
		}
		// enqueue the struct field
		sub := structField{
			rStructField: f.Type,
			parentIdx:    idx,
			parentPath:   path,
		}
		*queue = append(*queue, sub)
	}
	return nil
}

func getFieldName(f reflect.StructField, tag string) string {
	fieldName := f.Tag.Get(tag)
	fieldName = strings.SplitN(fieldName, ",", 2)[0]
	// maybe tag is not set.
	if fieldName == "" && !f.Anonymous {
		fieldName = f.Name
	}
	return fieldName
}
