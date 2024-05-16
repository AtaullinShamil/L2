package main

import (
	"bytes"
	"os/exec"
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "Test case-default",
			input: []string{"go", "run", "task.go", "5", "fileForTest.txt"},
			want:  "fasdfasdfsadfsgag 5\n",
		},
		{
			name:  "Test case-A",
			input: []string{"go", "run", "task.go", "-A=2", "5", "fileForTest.txt"},
			want:  "fasdfasdfsadfsgag 5\nasdfszfasf cat 6\nasvasfasf 7\n",
		},
		{
			name:  "Test case-B",
			input: []string{"go", "run", "task.go", "-B=2", "5", "fileForTest.txt"},
			want:  "qrewfsafjsadfj 3\nfsdfsdfsafasdfsd 4\nfasdfasdfsadfsgag 5\n",
		},
		{
			name:  "Test case-C",
			input: []string{"go", "run", "task.go", "-C=2", "5", "fileForTest.txt"},
			want:  "qrewfsafjsadfj 3\nfsdfsdfsafasdfsd 4\nfasdfasdfsadfsgag 5\nasdfszfasf cat 6\nasvasfasf 7\n",
		},
		{
			name:  "Test case-c",
			input: []string{"go", "run", "task.go", "-c=TRUE", "5", "fileForTest.txt"},
			want:  "1\n",
		},
		{
			name:  "Test case-i",
			input: []string{"go", "run", "task.go", "-i=TRUE", "CAT", "fileForTest.txt"},
			want:  "asdfszfasf cat 6\n",
		},
		{
			name:  "Test case-v",
			input: []string{"go", "run", "task.go", "-v=TRUE", "5", "fileForTest.txt"},
			want:  "abc 1\ndef 2\nqrewfsafjsadfj 3\nfsdfsdfsafasdfsd 4\nasdfszfasf cat 6\nasvasfasf 7\nasfasdf 8\nasdfsfsaasdf 9\nasfasd\n",
		},
		{
			name:  "Test case-F",
			input: []string{"go", "run", "task.go", "-F=asfasd", "5", "fileForTest.txt"},
			want:  "asfasd\n",
		},
		{
			name:  "Test case-n",
			input: []string{"go", "run", "task.go", "-n=TRUE", "5", "fileForTest.txt"},
			want:  "5: fasdfasdfsadfsgag 5\n",
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
				t.Errorf("Expected %v, got %v", tc.want, out.String())
			}
		})
	}
}
