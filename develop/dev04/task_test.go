package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramGroups(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  map[string][]string
	}{
		{
			name:  "Test case 1",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "Test case 2",
			input: []string{"вайд", "берри", "нос", "мышка", "голос", "столик"},
			want:  map[string][]string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findAnagramGroups(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}
