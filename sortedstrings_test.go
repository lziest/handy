package handy

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type testCase struct {
		ss            []string
		word          string
		expectedIndex int
	}

	testcases := []testCase{
		{
			ss:            []string{"a", "b", "c"},
			word:          "a",
			expectedIndex: 0,
		},
		{
			ss:            []string{"a", "b", "c"},
			word:          "b",
			expectedIndex: 1,
		},
		{
			ss:            []string{"a", "b", "c"},
			word:          "c",
			expectedIndex: 2,
		},
		{
			ss:            []string{"a", "b", "c"},
			word:          "d",
			expectedIndex: 3,
		},
		{
			ss:            []string{"c", "d"},
			word:          "a",
			expectedIndex: 0,
		},
		{
			ss:            []string{"c", "d"},
			word:          "b",
			expectedIndex: 0,
		},
		{
			ss:            []string{"c", "d"},
			word:          "c",
			expectedIndex: 0,
		},
		{
			ss:            []string{"c", "d"},
			word:          "d",
			expectedIndex: 1,
		},
		{
			ss:            []string{"c", "d"},
			word:          "e",
			expectedIndex: 2,
		},
		{
			ss:            []string{"c", "y"},
			word:          "x",
			expectedIndex: 1,
		},
		{
			ss:            []string{"c", "y"},
			word:          "y",
			expectedIndex: 1,
		},
		{
			ss:            []string{"c", "y"},
			word:          "z",
			expectedIndex: 2,
		},
		{
			ss:            []string{"c"},
			word:          "z",
			expectedIndex: 1,
		},
		{
			ss:            []string{"c"},
			word:          "a",
			expectedIndex: 0,
		},
		{
			ss:            []string{"c"},
			word:          "c",
			expectedIndex: 0,
		},
	}

	for _, c := range testcases {
		ss := SortedStrings(c.ss)
		ret := ss.BinarySearch(c.word)
		if ret != c.expectedIndex {
			t.Errorf("unexpected result %d, %v", ret, c)
		}
	}
}

func TestMinus(t *testing.T) {
	type testCase struct {
		ss, yass, expected []string
	}

	testcases := []testCase{
		{
			ss:       []string{"a", "b", "c"},
			yass:     []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			ss:       []string{"a", "b", "c"},
			yass:     []string{"b"},
			expected: []string{"a", "c"},
		},
		{
			ss:       []string{"a", "b", "c"},
			yass:     []string{"b", "c"},
			expected: []string{"a"},
		},
		{
			ss:       []string{"a", "b", "c"},
			yass:     []string{"b", "d"},
			expected: []string{"a", "c"},
		},
		{
			ss:       []string{"a", "b", "c", "z"},
			yass:     []string{"b", "d", "z"},
			expected: []string{"a", "c"},
		},
	}

	for _, c := range testcases {
		ss := SortedStrings(c.ss)
		ret := ss.Minus(c.yass)
		if len(ret) != len(c.expected) {
			t.Fatalf("unexpected result len %d, expected %d, %v", len(ret), len(c.expected), c)
		}
		for i, _ := range c.expected {
			if c.expected[i] != ret[i] {
				t.Errorf("unexpected result %v in index %d,  %v", ret, i, c.expected)
			}
		}
	}
}
