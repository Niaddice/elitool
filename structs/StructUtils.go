package structs

import (
	"reflect"
)

// 结构体相关工具类

// Struct2map
// @param resourceStruct 需要转换的结构体
// @param keyField 主键名称,结构体字段名称
// @param valueField 指定值的字段名称,当value是"" 空字符时,返回的map为 map[key]resource
func Struct2map(resourceStruct interface{}, keyField string, valueField string) map[string]interface{} {
	result := make(map[string]interface{})
	if resourceStruct == nil {
		panic("结构体不能为nil")
	}
	if valueField == "" {
		t := reflect.ValueOf(resourceStruct)
		if t.Kind() == reflect.Struct {
			keyValue := reflect.ValueOf(resourceStruct).FieldByName(keyField).String()
			result[keyValue] = nil
			if valueField != "" {
				value := reflect.ValueOf(resourceStruct).FieldByName(valueField).String()
				result[keyValue] = value
			} else {
				result[keyValue] = resourceStruct
			}
		}
	}
	return result
}
