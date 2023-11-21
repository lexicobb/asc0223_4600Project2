/* Author:      Alexia Cobb (alexiacobb@my.unt.edu)
 * EUID:		asc0223
 * Assignment:  Project 2 - Shell Builtins
 * Class:       CSCE 3550
 * Instructor:  Dr. Hochstetler
 * Due Date:    10 December 2023

 * File: main_test.go
 */

package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_runLoop(t *testing.T) {
	t.Parallel()

	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "exit command",
			args: args{
				r: strings.NewReader("exit\n"),
			},
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
		{
			name: "handleInput error",
			args: args{
				r: strings.NewReader("unknownCommand\nexit\n"),
			},
			wantErrW: "unknownCommand",
		},
		{
			name: "printPrompt error",
			args: args{
				r: &errorReader{},
			},
			wantErrW: "printPrompt error",
		},
		{
			name: "executeCommand error",
			args: args{
				r: strings.NewReader("cmdWithError\nexit\n"),
			},
			wantErrW: "cmdWithError",
		},
		{
			name: "no error with command",
			args: args{
				r: strings.NewReader("echo hello\nexit\n"),
			},
			wantW: "hello",
		},
		{
			name: "command not found",
			args: args{
				r: strings.NewReader("unknownCommand\nexit\n"),
			},
			wantErrW: "Command not found",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			exit := make(chan struct{}, 2)
			go runLoop(tt.args.r, w, errW, exit)
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			require.NotEmpty(t, w.String())
			if tt.wantW != "" {
				require.Contains(t, w.String(), tt.wantW)
			}
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			} else {
				require.Empty(t, errW.String())
			}
		})
	}
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, io.ErrUnexpectedEOF
}

func Test_printPrompt_Error(t *testing.T) {
	t.Parallel()

	w := &bytes.Buffer{}
	errW := &bytes.Buffer{}
	//errReader := &errorReader{}

	printPrompt(w) // Using printPrompt instead of printPromptError
	//printPromptError(errW, errReader)

	require.Empty(t, w.String())
	require.Contains(t, errW.String(), "printPrompt error")
}

func Test_executeCommand_Error(t *testing.T) {
	t.Parallel()

	w := &bytes.Buffer{}
	errW := &bytes.Buffer{}

	err := executeCommand(w.String(), "nonexistent-command")

	require.Error(t, err)
	require.Contains(t, err.Error(), "nonexistent-command")
	require.Contains(t, errW.String(), "nonexistent-command")
}

func Test_handleInput_CommandNotFound(t *testing.T) {
	t.Parallel()

	w := &bytes.Buffer{}
	//errW := &bytes.Buffer{}

	err := handleInput(w, "nonexistent-command", make(chan struct{}))

	require.Error(t, err)
	require.Contains(t, err.Error(), "Command not found")
}
