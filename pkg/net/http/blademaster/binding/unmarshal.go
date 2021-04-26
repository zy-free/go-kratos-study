package binding

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"
)

const (
	jsonTagName    = "json"
	defaultTagName = "default"
)


// DefaultWithNoInput set default value by tag if value is not input
func DefaultWithNoInput(data []byte, val interface{}) error {
	var m map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&m); err != nil {
		return err
	}
	decoder = json.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(val); err != nil {
		return err
	}
	setDefaultValueWithNoInput(m, val)
	return nil
}

// DefaultWithNoInput set default value by tag if value is empty
func DefaultWithEmpty(data []byte, val interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(val); err != nil {
		return err
	}
	setDefaultValueWithEmpty(val)
	return nil
}

func setDefaultValueWithNoInput(m map[string]interface{}, val interface{}) {
	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	t := reflect.TypeOf(val).Elem()
	v := reflect.ValueOf(val).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		// 没传值，且有default标签，则设置defaultVal
		if _, exists := field.Tag.Lookup(jsonTagName); !exists {
			continue
		}

		if _, ok := m[field.Tag.Get(jsonTagName)];ok{
			continue
		}

		if _, exists := field.Tag.Lookup(defaultTagName); !exists {
			continue
		}

		defVal := field.Tag.Get(defaultTagName)
		setValue(value, defVal)
	}
}

func setDefaultValueWithEmpty(val interface{}) {
	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	t := reflect.TypeOf(val).Elem()
	v := reflect.ValueOf(val).Elem()
	for i := 0; i < t.NumField(); i++ {
		if _, exists := t.Field(i).Tag.Lookup(defaultTagName); !exists {
			continue
		}

		if v.Field(i).IsZero() {
			tagVal := t.Field(i).Tag.Get(defaultTagName)
			setValue(v.Field(i), tagVal)
		}
	}
}

func setValue(v reflect.Value, defVal string) {
	switch v.Kind() {
	case reflect.String:
		v.SetString(defVal)
		break
	case reflect.Bool:
		tmp, err := strconv.ParseBool(defVal)
		if err != nil {
			panic(err.Error())
		}
		v.SetBool(tmp)
		break
	case reflect.Int64, reflect.Int32, reflect.Int, reflect.Int8:
		tmp, err := strconv.ParseInt(defVal, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		v.SetInt(tmp)
		break
	case reflect.Uint64, reflect.Uint32, reflect.Uint, reflect.Uint8:
		tmp, err := strconv.ParseUint(defVal, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		v.SetUint(tmp)
		break
	default:
		panic("unsupported type :" + v.Type().Kind().String())
	}
}
