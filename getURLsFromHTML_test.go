package main

import (
	"testing"
	"reflect"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name string
		inputURL string
		inputBody string
		expected []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name: "absolute only",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Boot.dev Blog</title>
			<link rel="stylesheet" href="https://blog.boot.dev/styles/main.css">
		</head>
		<body>
			<header>
				<a href="https://blog.boot.dev/images"><img src="https://blog.boot.dev/images/logo.png" alt="Logo"></a>
			</header>
			<main>
				<h1>Welcome to Boot.dev Blog</h1>
				<nav>
					<a href="https://blog.boot.dev/tutorials">Tutorials</a>
					<a href="https://blog.boot.dev/articles">Articles</a>
				</nav>
			</main>
			<footer>
				<a href="https://blog.boot.dev/privacy-policy">Privacy Policy</a>
			</footer>
			<script src="https://blog.boot.dev/scripts/main.js"></script>
		</body>
		</html>
			`,
			expected: []string{"https://blog.boot.dev/images", "https://blog.boot.dev/tutorials", "https://blog.boot.dev/articles", "https://blog.boot.dev/privacy-policy"},
		},
		{
			name: "relative only",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/1">first</a>
					<a href="/path/2">second</a>
					<a href="/path/3">third</a>
				</body>
			</html>
			`,
			expected: []string{"https://blog.boot.dev/path/1", "https://blog.boot.dev/path/2", "https://blog.boot.dev/path/3"},
		},
		{
			name: "empty",
			inputURL: "https://blog.boot.dev",
			inputBody: ``,
			expected: []string{},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "**&^blog()boot:dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected){
				t.Errorf("Test %v - %s FAIL: expected URLS: %v, actual: %v", i, tc.name, tc.expected, actual)
			} 
		})
	}
}
