package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	filename := "data.txt"
	data := []string{
		"John\tDoe\t25",
		"Jane\tSmith\t40",
		"Alice\tJohnson\t28",
		"Hloya-Bonen-11",
		"Izya-Hamul-75",
		"Jib-Beater-37",
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range data {
		_, _ = fmt.Fprintln(writer, line)
	}
	writer.Flush()

	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "Test case-f=1",
			input: []string{"go", "run", "task.go", "-f=1", "data.txt"},
			want:  "John\nJane\nAlice\nHloya-Bonen-11\nIzya-Hamul-75\nJib-Beater-37\n",
		},
		{
			name:  "Test case-f=1,2",
			input: []string{"go", "run", "task.go", "-f=1,2", "data.txt"},
			want:  "John\tDoe\nJane\tSmith\nAlice\tJohnson\nHloya-Bonen-11\nIzya-Hamul-75\nJib-Beater-37\n",
		},
		{
			name:  "Test case-d",
			input: []string{"go", "run", "task.go", "-f=1", "-d=-", "data.txt"},
			want:  "John\tDoe\t25\nJane\tSmith\t40\nAlice\tJohnson\t28\nHloya\nIzya\nJib\n",
		},
		{
			name:  "Test case-s",
			input: []string{"go", "run", "task.go", "-f=1", "-s=TRUE", "data.txt"},
			want:  "John\nJane\nAlice\n",
		},
		{
			name:  "Test case-s,d",
			input: []string{"go", "run", "task.go", "-f=1", "-s=TRUE", "-d=-", "data.txt"},
			want:  "Hloya\nIzya\nJib\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command(tc.input[0], tc.input[1:]...)

			var out bytes.Buffer
			cmd.Stdout = &out
			if err := cmd.Run(); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(out.String(), tc.want) {
				t.Errorf("Expected \n%v, got \n%v", tc.want, out.String())
			}
		})
	}
}
