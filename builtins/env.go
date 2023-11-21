/* Author:      Alexia Cobb (alexiacobb@my.unt.edu)
 * EUID:		asc0223
 * Assignment:  Project 2 - Shell Builtins
 * Class:       CSCE 3550
 * Instructor:  Dr. Hochstetler
 * Due Date:    10 December 2023

 * File: env.go
 * already provided by Dr. Hochstetler
 */

package builtins

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func EnvironmentVariables(w io.Writer, args ...string) error {
	toRemove := make([]string, 0)
	for i := 0; i < len(args); i++ {
		if args[i] == "-u" {
			if len(args) < i+2 {
				return fmt.Errorf("%w: -u requires an argument", ErrInvalidArgCount)
			}
			toRemove = append(toRemove, args[i+1])
			i++
		}
	}

	toShow := make([]string, 0)
	for _, env := range os.Environ() {
		show := true
		for _, v := range toRemove {
			if strings.HasPrefix(env, v+"=") {
				show = false
				break
			}
		}
		if show {
			toShow = append(toShow, env)
		}
	}

	_, err := fmt.Fprintln(w, strings.Join(toShow, "\n"))

	return err
}
