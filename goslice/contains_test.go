package goslice

import "testing"

// Dummy is a test dummy struct
type Dummy struct {
	Age  int
	Name string
}

// DummyV2 is another dummy test struct
type DummyV2 struct {
	Address       string
	IsSuperSaiyan bool
}

// Simple comparison function which always returns false
var alwaysNotEqual = func(a, b interface{}) bool {
	return false
}

// Simple equals function for Dummy structs
var dummyEquals = func(a, b interface{}) bool {
	convertedA, _ := a.(Dummy)
	convertedB, _ := b.(Dummy)
	return convertedA.Age == convertedB.Age && convertedA.Name == convertedB.Name
}

func TestStringInSlice(t *testing.T) {
	var tt = []struct {
		desc     string
		slice    []string
		entry    string
		contains bool
	}{
		// False
		{"empty string", []string{"Hi", "Hello", "Servus"}, "", false},
		{"not in slice basic", []string{"Hi", "Hello", "Servus"}, "Moin", false},
		{"check empty slice", []string{}, "Servus", false},

		// True
		{"basic successful", []string{"Hi", "Hello", "Servus"}, "Servus", true},
		{"contains empty string", []string{"", "Hello", "Servus"}, "", true},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			_, doesContain := StringInSlice(tc.slice, tc.entry)
			if doesContain != tc.contains {
				t.Errorf("Did not get expected result, was '%t', expected '%t' with input '%v' and '%s'", doesContain, tc.contains, tc.slice, tc.entry)
			}
		})
	}
}

func TestIntInSlice(t *testing.T) {
	var tt = []struct {
		desc     string
		slice    []int
		entry    int
		contains bool
	}{
		// False
		{"basic unsuccessful", []int{1, 2, 3}, 4, false},
		{"check empty slice", []int{}, 1337, false},

		// True
		{"basic successful", []int{2, 12424, 453, 982395}, 453, true},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			_, doesContain := IntInSlice(tc.slice, tc.entry)
			if doesContain != tc.contains {
				t.Errorf("Did not get expected result, was '%t', expected '%t' with input '%v' and '%d'", doesContain, tc.contains, tc.slice, tc.entry)
			}
		})
	}
}

func TestBoolInSlice(t *testing.T) {
	var tt = []struct {
		desc     string
		slice    []bool
		entry    bool
		contains bool
	}{
		// False
		{"basic unsuccessful", []bool{false, false, false}, true, false},
		{"check empty slice", []bool{}, false, false},

		// True
		{"basic successful", []bool{true, false, true, true, true}, false, true},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			_, doesContain := BoolInSlice(tc.slice, tc.entry)
			if doesContain != tc.contains {
				t.Errorf("Did not get expected result, was '%t', expected '%t' with input '%v' and '%t'", doesContain, tc.contains, tc.slice, tc.entry)
			}
		})
	}
}

func TestStructInSlice(t *testing.T) {
	var tt = []struct {
		desc      string
		slice     interface{}
		entry     interface{}
		compFunc  ComparisonFunc
		contains  bool
		expectErr bool
	}{
		// False
		{"basic unsuccessful", []Dummy{Dummy{Age: 1234, Name: "Goku"}, Dummy{Age: 5678, Name: "Vegeta"}}, Dummy{Age: 9876, Name: "Trunks"}, nil, false, false},
		{"check empty slice", []Dummy{}, Dummy{Age: 12, Name: "Goku"}, nil, false, false},
		{"check wrong input slice", 1234, Dummy{Age: 12, Name: "Goku"}, nil, false, true},
		{"check different types slice", []Dummy{Dummy{Age: 12, Name: "Goku"}}, DummyV2{Address: "address", IsSuperSaiyan: false}, nil, false, false},
		{"check always not equal types slice", []Dummy{Dummy{Age: 12, Name: "Goku"}}, Dummy{Age: 12, Name: "Goku"}, alwaysNotEqual, false, false},

		// True
		{"basic successful", []Dummy{Dummy{Age: 12, Name: "Goku"}}, Dummy{Age: 12, Name: "Goku"}, nil, true, false},
		{"basic successful with function", []Dummy{Dummy{Age: 12, Name: "Goku"}}, Dummy{Age: 12, Name: "Goku"}, dummyEquals, true, false},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			_, doesContain, err := StructInSlice(tc.slice, tc.entry, tc.compFunc)
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error to occur for inputs '%v' and '%v', but got none %s", tc.slice, tc.entry, err)
			}
			if doesContain != tc.contains {
				t.Errorf("Did not get expected result, was '%t', expected '%t' with input '%v' and '%v'", doesContain, tc.contains, tc.slice, tc.entry)
			}
		})
	}
}
