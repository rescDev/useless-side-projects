package goslice

import (
	"reflect"
)

// ComparisonFunc declares a function type used for comparing custom Golang types
type ComparisonFunc func(a, b interface{}) bool

// StringInSlice checks if a string slice contains a given string and also returns its index
func StringInSlice(slice []string, entry string) (int, bool) {
	return checkIfElementExists(slice, entry)
}

// IntInSlice checks if an integer slice contains a given integer and also returns its index
func IntInSlice(slice []int, entry int) (int, bool) {
	return checkIfElementExists(slice, entry)
}

// BoolInSlice checks if a bool slice contains a given bool and also returns its index
func BoolInSlice(slice []bool, entry bool) (int, bool) {
	return checkIfElementExists(slice, entry)
}

// StructInSlice can be used for checking if a complex structs is in a slice of structs of the same type
func StructInSlice(slice interface{}, element interface{}, compareFunc ComparisonFunc) (int, bool, error) {
	convertedSlice, err := convertToSlice(slice)
	if err != nil {
		return -1, false, err
	}

	// If no custom func is provided, check for equality via reflect.DeepEqual
	if compareFunc == nil {
		compareFunc = func(a, b interface{}) bool {
			return reflect.DeepEqual(a, b)
		}
	}

	for i, s := range convertedSlice {
		if compareFunc(s, element) {
			return i, true, nil
		}
	}
	return -1, false, nil
}

// Helper func, checks if "slice" contains "element"
func checkIfElementExists(slice interface{}, element interface{}) (int, bool) {
	convertedSlice, _ := convertToSlice(slice)

	for i, s := range convertedSlice {
		if s == element {
			return i, true
		}
	}

	return -1, false
}
