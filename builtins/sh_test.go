/* Author:      Alexia Cobb (alexiacobb@my.unt.edu)
 * EUID:		asc0223
 * Assignment:  Project 2 - Shell Builtins
 * Class:       CSCE 3550
 * Instructor:  Dr. Hochstetler
 * Due Date:    10 December 2023

 * File: sh_test.go
 */

package builtins_test

import (
	"errors"
	"os/exec"
	"syscall"
	"testing"

	"github.com/lexicobb/asc0223_4600Project2/builtins"
)

func TestRunShellCommand(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name         string
		args         args
		wantErr      error
		wantExitCode int
	}{
		{
			name: "valid echo command",
			args: args{
				args: []string{"echo", "Hello, World!"},
			},
			wantErr:      nil,
			wantExitCode: 0,
		},
		{
			name: "invalid command",
			args: args{
				args: []string{"nonexistent-command"},
			},
			wantErr:      errors.New("command failed"),
			wantExitCode: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// testing
			err := builtins.Shell(tt.args.args...)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("Shell() expected error, but got nil")
				} else if err.Error() != tt.wantErr.Error() {
					t.Fatalf("Shell() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Fatalf("Shell() unexpected error: %v", err)
				}
			}

			// Check the exit code of the command (valid for Unix-like systems)
			exitCode := 0
			if err != nil {
				exitError, ok := err.(*exec.ExitError)
				if ok {
					exitStatus := exitError.Sys().(syscall.WaitStatus)
					exitCode = exitStatus.ExitStatus()
				}
			}
			if exitCode != tt.wantExitCode {
				t.Errorf("Shell() exit code = %d, wantExitCode %d", exitCode, tt.wantExitCode)
			}
		})
	}
}
