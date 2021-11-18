package cli

import (
	"fmt"
	"os"

	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/Octy-ai/octy-cli/pkg/output"
	"github.com/briandowns/spinner"
)

// Error type interface
type iError interface {
	outputError()
}

type Error struct {
	errorMsg     string
	extendedHelp string
	exitCode     int
	spinner      *spinner.Spinner
}

func (e Error) outputError() {
	if e.spinner != nil {
		output.StopSpinner(e.spinner, "\n", e.exitCode, os.Stdout)
	}
	mes := fmt.Sprintf("error: %s -- extended help: %s", e.errorMsg, e.extendedHelp)
	output.FPrint(mes)
}

// Error types

type unknownError struct {
	Error
}

func newUnknownError(msg string, spinner *spinner.Spinner) iError {
	return &unknownError{
		Error: Error{
			errorMsg:     msg,
			extendedHelp: fmt.Sprintf("Please try again. If this error persists, please open a support ticket at %s", globals.SupportTicketURL),
			exitCode:     1,
			spinner:      spinner,
		},
	}
}

type unauthorizedError struct {
	Error
}

func newUnauthorizedError(msg string, spinner *spinner.Spinner) iError {
	return &unauthorizedError{
		Error: Error{
			errorMsg:     msg,
			extendedHelp: "Please update your Octy credentials by running the command 'octy-cli auth -p -s' providing valid Octy public and secret keys",
			exitCode:     1,
			spinner:      spinner,
		},
	}
}

type badRequestError struct {
	Error
}

func newBadRequestError(msg string, spinner *spinner.Spinner) iError {
	return &badRequestError{
		Error: Error{
			errorMsg:     msg,
			extendedHelp: "n/a",
			exitCode:     1,
			spinner:      spinner,
		},
	}
}

type unprocessableEntityError struct {
	Error
}

func newUnprocessableEntityError(msg string, spinner *spinner.Spinner) iError {
	return &unprocessableEntityError{
		Error: Error{
			errorMsg:     msg,
			extendedHelp: fmt.Sprintf("Refer to the documentation %s", globals.Docs),
			exitCode:     1,
			spinner:      spinner,
		},
	}
}
