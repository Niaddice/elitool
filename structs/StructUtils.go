package structs

import (
	"reflect"
)

// 结构体相关工具类

// Struct2Map 结构体转map
// @param resourceStruct 需要转换的结构体
// @param keyField 主键名称,结构体字段名称
// @param valueField 指定值的字段名称,当value是"" 空字符时,返回的map为 map[key]resource
func Struct2Map(resourceStruct interface{}, keyField string, valueField string) map[string]interface{} {
	result := make(map[string]interface{})
	key, value := struct2KV(resourceStruct, keyField, valueField)
	result[key] = value
	return result
}

// StructSlice2Map 将结构体切片转成map
// @param resourceStruct 切片
// @param keyField 主键名称,结构体字段名称
// @param valueField 指定值的字段名称,当value是"" 空字符时,返回的map为 map[key]resource
func StructSlice2Map(resourceStruct interface{}, keyField string, valueField string) map[string]interface{} {
	result := make(map[string]interface{})
	// 判断传入的resourceStruct是否是切片
	val := reflect.ValueOf(resourceStruct)
	if val.Kind() == reflect.Slice {
		length := val.Len()
		for i := 0; i < length; i++ {
			v := val.Index(i).Interface()
			key, value := struct2KV(v, keyField, valueField)
			result[key] = value
		}
	} else {
		panic("resourceStruct必须是切片!")
	}
	return result
}

// Struct2KV
// @param resourceStruct 切片
// @param keyField 主键名称,结构体字段名称
// @param valueField 指定值的字段名称,当value是"" 空字符时,返回的map为 map[key]resource
func struct2KV(resourceStruct interface{}, keyField string, valueField string) (string, interface{}) {
	var key string
	var value interface{}
	if resourceStruct == nil {
		panic("结构体不能为nil")
	}
	t := reflect.ValueOf(resourceStruct)
	if t.Kind() == reflect.Struct {
		key = reflect.ValueOf(resourceStruct).FieldByName(keyField).String()
		if valueField != "" {
			value = reflect.ValueOf(resourceStruct).FieldByName(valueField)
		} else {
			value = resourceStruct
		}
	}
	return key, value
}
