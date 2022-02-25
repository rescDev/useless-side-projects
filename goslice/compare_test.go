package goslice

import "testing"

func TestCompareStringSlices(t *testing.T) {
	var tt = []struct {
		desc   string
		slice1 []string
		slice2 []string
		equal  bool
	}{
		// False
		{"empty slice", []string{"Hi", "Hello", "Servus"}, []string{}, false},
		{"basic not equal", []string{"Hi", "Hello", "Servus"}, []string{"Moin"}, false},
		{"wrong order not equal", []string{"Hi", "Hello", "Servus"}, []string{"Hi", "Servus", "Hello"}, false},

		// True
		{"basic equal", []string{"Hi", "Hello", "Servus"}, []string{"Hi", "Hello", "Servus"}, true},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			isEqual := CompareStringSlices(tc.slice1, tc.slice2)
			if isEqual != tc.equal {
				t.Errorf("Did not get expected result, was '%t', expected '%t' with input '%v' and '%v'", isEqual, tc.equal, tc.slice1, tc.slice2)
			}
		})
	}
}

func TestCompareIntSlices(t *testing.T) {
	var tt = []struct {
		desc   string
		slice1 []int
		slice2 []int
		equal  bool
	}{
		// False
		{"empty slice", []int{1, 2, 3}, []int{}, false},
		{"basic not equal", []int{1, 2, 3}, []int{2, 4}, false},
		{"wrong order not equal", []int{1, 2, 3}, []int{3, 2, 1}, false},

		// True
		{"basic equal", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, true},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			isEqual := CompareIntSlices(tc.slice1, tc.slice2)
			if isEqual != tc.equal {
				t.Errorf("Did not get expected result, was '%t', expected '%t' with input '%v' and '%v'", isEqual, tc.equal, tc.slice1, tc.slice2)
			}
		})
	}
}

func TestCompareBoolSlices(t *testing.T) {
	var tt = []struct {
		desc   string
		slice1 []bool
		slice2 []bool
		equal  bool
	}{
		// False
		{"empty slice", []bool{true, false, true}, []bool{}, false},
		{"basic not equal", []bool{true, false}, []bool{false}, false},
		{"wrong order not equal", []bool{true, false}, []bool{false, true}, false},

		// True
		{"basic equal", []bool{true, true, false}, []bool{true, true, false}, true},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			isEqual := CompareBoolSlices(tc.slice1, tc.slice2)
			if isEqual != tc.equal {
				t.Errorf("Did not get expected result, was '%t', expected '%t' with input '%v' and '%v'", isEqual, tc.equal, tc.slice1, tc.slice2)
			}
		})
	}
}

func TestCompareStructSlices(t *testing.T) {
	var tt = []struct {
		desc        string
		slice1      interface{}
		slice2      interface{}
		compareFunc ComparisonFunc
		equal       bool
	}{
		// False
		{"empty slice", []Dummy{Dummy{Age: 1234, Name: "Goku"}}, nil, nil, false},
		{"basic not equal", []Dummy{Dummy{Age: 1234, Name: "Goku"}, Dummy{Age: 4321, Name: "Vegeta"}}, []Dummy{Dummy{Age: 4321, Name: "Vegeta"}}, nil, false},
		{"wrong order not equal", []Dummy{Dummy{Age: 1234, Name: "Goku"}, Dummy{Age: 4321, Name: "Vegeta"}}, []Dummy{Dummy{Age: 4321, Name: "Vegeta"}, Dummy{Age: 1234, Name: "Goku"}}, nil, false},

		// True
		{"basic equal", []Dummy{Dummy{Age: 1234, Name: "Goku"}}, []Dummy{Dummy{Age: 1234, Name: "Goku"}}, nil, true},
		{"basic equal with func", []Dummy{Dummy{Age: 1234, Name: "Goku"}}, []Dummy{Dummy{Age: 1234, Name: "Goku"}}, dummyEquals, true},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			isEqual, _ := CompareStructSlices(tc.slice1, tc.slice2, tc.compareFunc)
			if isEqual != tc.equal {
				t.Errorf("Did not get expected result, was '%t', expected '%t' with input '%v' and '%v'", isEqual, tc.equal, tc.slice1, tc.slice2)
			}
		})
	}
}
