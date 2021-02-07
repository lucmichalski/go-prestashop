package struct2map

import (
	"errors"
	"reflect"
	"strings"
)

var (
	ErrNotPtr       = errors.New("need a pointer")
	ErrNotValidElem = errors.New("pointer not point to struct")
	ErrNotValidTag  = errors.New("not valid tag")
	ErrNotValidKey  = errors.New("not valid key")
	ErrIgnore       = errors.New("ignore key")
	ErrNeedTag      = errors.New("need struct2map tag")
	TagName         = "struct2map"
	TagKey          = "key"
	TagIgnore       = "-"
)

func getKey(tagStr string) (key string, err error) {
	err = ErrNotValidTag
	for _, tag := range strings.Split(tagStr, ";") {
		kv := strings.Split(tag, ":")

		if kv[0] == TagKey {
			if kv[1] == "" {
				err = ErrNotValidKey
			} else {
				key = kv[1]
				err = nil
			}
		} else if kv[0] == TagIgnore {
			err = ErrIgnore
		}
	}

	return key, err
}

func Decode(st interface{}) (map[string]interface{}, error) {
	stType := reflect.TypeOf(st)
	if stType.Kind() != reflect.Ptr {
		return nil, ErrNotPtr
	}

	eleType := stType.Elem()
	if eleType.Kind() != reflect.Struct {
		return nil, ErrNotValidElem
	}

	stVal := reflect.Indirect(reflect.ValueOf(st))
	m := make(map[string]interface{})

	for i := 0; i < eleType.NumField(); i++ {
		tagStr, ok := eleType.Field(i).Tag.Lookup(TagName)
		if !ok {
			return nil, ErrNeedTag
		}

		key, err := getKey(tagStr)
		if err == ErrNotValidKey || err == ErrNotValidTag {
			return nil, err
		}

		if err == ErrIgnore {
			continue
		}

		m[key] = stVal.Field(i).Interface()
	}

	return m, nil
}
