package slice

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/xiagoo/gosort/consts"
)

type baseSort struct {
	length int
	less   func(i, j int) bool
	swap   func(i, j int)
}

func (bs *baseSort) Len() int           { return bs.length }
func (bs *baseSort) Less(i, j int) bool { return bs.less(i, j) }
func (bs *baseSort) Swap(i, j int)      { bs.swap(i, j) }

//SortByKey sort slice by key, key should by struct field name, sort type contain asc and desc
func SortByKey(slice interface{}, sortKey string, sortType int) {
	st := reflect.TypeOf(slice)

	if st.Kind() != reflect.Slice {
		panic(fmt.Sprintf("type must be slice , you type is %s", st.Kind()))
	}

	if st.Elem().Kind() != reflect.Ptr {
		panic(fmt.Sprintf("type must be ptr, you type is %s", st.Elem().Kind()))
	}

	if st.Elem().Elem().Kind() != reflect.Struct {
		panic(fmt.Sprintf("type must be struct, you type is %s", st.Elem().Elem().Kind()))
	}

	field, flag := st.Elem().Elem().FieldByName(sortKey)
	if !flag {
		panic(fmt.Sprintf("struct field doesn't exist %s", sortKey))
	}

	sv := reflect.ValueOf(slice)

	less := func(i, j int) bool {
		svi := sv.Index(i).Elem().FieldByName(sortKey)
		svj := sv.Index(j).Elem().FieldByName(sortKey)
		switch field.Type.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if consts.Asc == sortType {
				return svi.Int() < svj.Int()
			}
			return svi.Int() > svj.Int()
		case reflect.Float32, reflect.Float64:
			if consts.Asc == sortType {
				return svi.Float() < svj.Float()
			}
			return svi.Float() > svj.Float()
		case reflect.String:
			if consts.Asc == sortType {
				return svi.String() < svj.String()
			}
			return svi.String() > svj.String()
		default:
			return false
		}
	}

	sort.Sort(sortSlice(slice, less))
}

func sortSlice(slice interface{}, less func(i, j int) bool) sort.Interface {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("slice.Sort need slice value of type %T", slice))
	}
	return &baseSort{
		length: sv.Len(),
		less:   less,
		swap:   swapper(sv),
	}
}

func swapper(v reflect.Value) func(i, j int) {
	tmp := reflect.New(v.Type().Elem()).Elem()
	return func(i, j int) {
		v1 := v.Index(i)
		v2 := v.Index(j)
		tmp.Set(v1)
		v1.Set(v2)
		v2.Set(tmp)
	}
}
