package stringsheet

import (
	"reflect"
	"testing"
)

func TestComputeKMPLPS(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		exp    []int
		expErr error
	}{
		{"empty pattern", "", nil, EmptyPatternErr},
		{"small pattern", "aabaab", []int{0, 1, 0, 1, 2, 3}, nil},
		{"repetitive pattern", "AAAA", []int{0, 1, 2, 3}, nil},
		{"different pattern", "ABCDE", []int{0, 0, 0, 0, 0}, nil},
		{"uncommon pattern", "AABAACAABAA", []int{0, 1, 0, 1, 2, 0, 1, 2, 3, 4, 5}, nil},
		{"normal pattern", "AAACAAAAAC", []int{0, 1, 2, 0, 1, 2, 3, 3, 3, 4}, nil},
		{"typical pattern", "AAABAAA", []int{0, 1, 2, 0, 1, 2, 3}, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ComputeKMPLPS(tc.input)
			// expected error path
			if tc.expErr != nil {
				if err == nil {
					t.Errorf("expected error but got nil instead")
				}
				if err != tc.expErr {
					t.Errorf("expected error %v, but got %v instead", tc.expErr, err)
				}
				if actual != nil {
					t.Errorf("expected nil result, but got %v instead", actual)
				}
				return
			}

			if err != nil {
				t.Errorf("expected nil error, but got %v instead", err)
			}
			if !reflect.DeepEqual(actual, tc.exp) {
				t.Errorf("expected %v, but got %v instead", tc.exp, actual)
			}
		})
	}
}

func TestKMPMatch(t *testing.T) {
	testCases := []struct {
		name         string
		inputPattern string
		inputText    string
		exp          []int
		expErr       error
	}{
		{"empty pattern", "", "tst", nil, EmptyPatternErr},
		{"empty text", "test", "", nil, EmptyTextErr},
		{"match found 1", "aabaab", "aabaabaab", []int{0, 3}, nil},
		{"match found 2", "ABABCABAB", "ABABDABACDABABCABAB", []int{10}, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := KMPMatch(tc.inputPattern, tc.inputText)
			if tc.expErr != nil {
				if err == nil {
					t.Errorf("expected error %v, but got nil instead", tc.expErr)
				}
				if err != tc.expErr {
					t.Errorf("expected error %v, but got %v instead", tc.expErr, err)
				}
				if actual != nil {
					t.Errorf("expected nil, but got %v instead", actual)
				}
				return
			}
			if err != nil {
				t.Errorf("expected nil error, but got %v instead", err)
			}
			if !reflect.DeepEqual(actual, tc.exp) {
				t.Errorf("expected %v, but got %v instead", tc.exp, actual)
			}
		})
	}
}
