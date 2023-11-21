/* Author:      Alexia Cobb (alexiacobb@my.unt.edu)
 * EUID:         asc0223
 * Assignment:  Project 2 - Shell Builtins
 * Class:       CSCE 3550
 * Instructor:  Dr. Hochstetler
 * Due Date:    10 December 2023

 * File: bash.go
 */

package builtins

import (
	"fmt"
	"os"
	"os/exec"
)

func Bash(args ...string) error {
	// Check if there are any arguments provided to the bash command.
	if len(args) == 0 {
		return fmt.Errorf("bash: missing argument")
	}

	// Execute the bash command with the provided arguments.
	cmd := exec.Command("bash", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("bash: %v", err)
	}

	return nil
}
