package main

import (
	"errors"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
			err:      nil,
		},
		{
			name:     "One character",
			input:    "a",
			expected: "a",
			err:      nil,
		},
		{
			name:     "String starting with number",
			input:    "3a",
			expected: "",
			err:      errors.New("invalid string"),
		},
		{
			name:     "Normal string",
			input:    "ab2c3",
			expected: "abbccc",
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decode(tt.input)
			if (err != nil) != (tt.err != nil) {
				t.Errorf("decode() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.expected {
				t.Errorf("decode() = %v, want %v", got, tt.expected)
			}
		})
	}
}
