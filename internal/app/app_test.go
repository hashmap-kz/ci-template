package app

import "testing"

func TestIsValidIdentifier(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"valid identifier", "myVariable", true},
		{"starts with underscore", "_privateVar", true},
		{"contains digits", "var123", true},
		{"contains only underscore", "_", true},
		{"starts with digit", "123var", false},
		{"contains space", "my var", false},
		{"contains special char", "var@name", false},
		{"just a number", "123", false},
		{"single letter", "a", true},
		{"unicode letter", "变量", true},
		{"leading underscore with digits", "_123abc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidIdentifier(tt.input)
			if result != tt.expected {
				t.Errorf("IsValidIdentifier(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
