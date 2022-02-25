package goslice

import "reflect"

// CompareStringSlices compares two string slices
func CompareStringSlices(s1, s2 []string) bool {
	return checkEquality(s1, s2)
}

// CompareIntSlices compares two int slices
func CompareIntSlices(s1, s2 []int) bool {
	return checkEquality(s1, s2)
}

// CompareBoolSlices compares two bool slices
func CompareBoolSlices(s1, s2 []bool) bool {
	return checkEquality(s1, s2)
}

// CompareStructSlices compares two struct slices
func CompareStructSlices(s1, s2 interface{}, compareFunc ComparisonFunc) (bool, error) {
	convertedS1, err := convertToSlice(s1)
	if err != nil {
		return false, err
	}

	convertedS2, err := convertToSlice(s2)
	if err != nil {
		return false, err
	}

	if (convertedS1 == nil) != (convertedS2 == nil) {
		return false, nil
	}

	if len(convertedS1) != len(convertedS2) {
		return false, nil
	}

	if compareFunc == nil {
		compareFunc = func(a, b interface{}) bool {
			return reflect.DeepEqual(a, b)
		}
	}

	// check should be ok since it is already checked that slices have the same length
	for i, v := range convertedS1 {
		if !compareFunc(v, convertedS2[i]) {
			return false, nil
		}
	}

	return true, nil
}

// Helper function to check for equality
func checkEquality(s1, s2 interface{}) bool {
	convertedS1, _ := convertToSlice(s1)
	convertedS2, _ := convertToSlice(s2)

	if (convertedS1 == nil) != (convertedS2 == nil) {
		return false
	}

	if len(convertedS1) != len(convertedS2) {
		return false
	}

	// check should be ok since it is already checked that slices have the same length
	for i, v := range convertedS1 {
		if v != convertedS2[i] {
			return false
		}
	}

	return true
}
