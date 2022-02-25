package goslice

import (
	"fmt"
	"reflect"
)

// converts interface{} to []interface{}
func convertToSlice(slice interface{}) ([]interface{}, error) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil, fmt.Errorf("submitted slice is from non-slice type '%s', convertion cancelled", s.Kind().String())
	}

	result := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		result[i] = s.Index(i).Interface()
	}

	return result, nil
}
