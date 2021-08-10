package structs

import "reflect"

func struct2map(resource interface{}) map[string]interface{} {
	typeOf := reflect.TypeOf(resource)
	typeOf.Name()
	return nil
}
