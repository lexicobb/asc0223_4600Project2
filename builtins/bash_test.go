/* Author:      Alexia Cobb (alexiacobb@my.unt.edu)
 * EUID:		asc0223
 * Assignment:  Project 2 - Shell Builtins
 * Class:       CSCE 3550
 * Instructor:  Dr. Hochstetler
 * Due Date:    10 December 2023

 * File: bash_test.go
 */

package builtins_test

import (
	"errors"
	"testing"

	"github.com/lexicobb/asc0223_4600Project2/builtins"
)

func TestBashCommand(t *testing.T) {
	//tmp := t.TempDir()

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "error no args",
			args:    args{},
			wantErr: builtins.ErrInvalidArgCount,
		},
		{
			name: "run bash command",
			args: args{
				args: []string{"echo Hello, World!"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// testing
			if err := builtins.Bash(tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("Bash() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("Bash() unexpected error: %v", err)
			}

			// additional checks if needed
		})
	}
}
