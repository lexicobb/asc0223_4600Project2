/* Author:      Alexia Cobb (alexiacobb@my.unt.edu)
 * EUID:		asc0223
 * Assignment:  Project 2 - Shell Builtins
 * Class:       CSCE 3550
 * Instructor:  Dr. Hochstetler
 * Due Date:    10 December 2023

 * File: ksh.go
 */

package builtins

import (
	"fmt"
	"os"
	"os/exec"
)

func Ksh() error {
	// Execute the ksh command.
	cmd := exec.Command("ksh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ksh: %v", err)
	}

	return nil
}
