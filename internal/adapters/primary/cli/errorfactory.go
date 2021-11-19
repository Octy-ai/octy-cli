package cli

import (
	"strings"
	"time"

	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/briandowns/spinner"
	"github.com/getsentry/sentry-go"
)

// errorFactory: handles outputting error messages with extended help
func errorFactory(errs []error, shouldQuit bool, spinner *spinner.Spinner) {

	errDuplicates := []string{}
	var e iError

ErrorLoop:
	for _, err := range errs {
		var prefix string
		var msg string = err.Error()

		// prevent duplicate errors being processed.
		for _, s := range errDuplicates {
			if s == err.Error() {
				continue ErrorLoop
			}
		}
		errDuplicates = append(errDuplicates, err.Error())

		if strings.Contains(err.Error(), ":") {
			sections := strings.Split(err.Error(), ":")
			prefix = sections[0]
			msg = sections[1]
		}

		switch prefix {
		case "apierror[500]":
			e = newUnknownError(msg, spinner)
		case "apierror[422]":
			e = newUnprocessableEntityError(msg, spinner)
		case "apierror[401]":
			e = newUnauthorizedError(msg, spinner)
		case "apierror[400]":
			e = newBadRequestError(msg, spinner)
		default:
			logException(err)
			e = newUnknownError(globals.ErrUnknownError.Error(), spinner)
		}
		e.outputError()
	}
	if shouldQuit {
		quit("", 1, nil)
	}

}

//
// Private functions
//

func logException(exception error) {
	sentry.CaptureException(exception)
	sentry.Flush(time.Second * 5)
}
