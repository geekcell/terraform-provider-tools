package utils

import "reflect"

func RemoveNullValues(m map[string]interface{}) {
	val := reflect.ValueOf(m)
	for _, e := range val.MapKeys() {
		v := val.MapIndex(e)
		if v.IsNil() {
			delete(m, e.String())
			continue
		}

		switch t := v.Interface().(type) {
		// If key is a map
		case map[string]interface{}:
			RemoveNullValues(t)

		// If key is an array
		case []interface{}:
			for _, s := range t {
				// Check if the values inside the array are JSON objects (Go Map)
				if s != nil && reflect.TypeOf(s).Kind() == reflect.Map {
					RemoveNullValues(s.(map[string]interface{}))
				}
			}
		}
	}
}
