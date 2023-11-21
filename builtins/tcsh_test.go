/* Author:      Alexia Cobb (alexiacobb@my.unt.edu)
 * EUID:		asc0223
 * Assignment:  Project 2 - Shell Builtins
 * Class:       CSCE 3550
 * Instructor:  Dr. Hochstetler
 * Due Date:    10 December 2023

 * File: tcsh_test.go
 */

package builtins

import (
	"io"
	"os"
	"testing"
)

func TestTcsh(t *testing.T) {
	// Capture the current standard input, output, and error streams.
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	oldStderr := os.Stderr

	// Redirect standard input, output, and error for testing.
	r, w, _ := os.Pipe()
	os.Stdin = r
	os.Stdout = w
	os.Stderr = w

	// Close the write end of the pipe when the test is done.
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
		os.Stderr = oldStderr
		w.Close()
	}()

	// Run the tcsh command.
	err := TCShell()
	if err != nil {
		t.Fatalf("Tcsh() error: %v", err)
	}

	// Close the read end of the pipe to stop reading from it.
	r.Close()

	// Read the captured output from the command.
	out, err := io.ReadAll(w)
	if err != nil {
		t.Fatalf("Error reading output from command: %v", err)
	}

	// Check if the output contains an expected string.
	expectedString := "some_expected_output_string"
	if !containsString(string(out), expectedString) {
		t.Errorf("Expected output to contain: %s\nActual output: %s", expectedString, string(out))
	}
}
