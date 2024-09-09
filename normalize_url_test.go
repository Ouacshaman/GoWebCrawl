package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name: "normalize path 01",
			inputURL: "http://api.boot.dev/path/path2",
			expected: "api.boot.dev/path/path2",
		},
		{
			name: "normalize path 01",
			inputURL: "http://boot.dev",
			expected: "boot.dev",
		},
		{
			name: "invalid slop",
			inputURL: "kek&&as::dfab.op",
			expected: "",
		},
		{
			name: "empty",
			inputURL: "",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
