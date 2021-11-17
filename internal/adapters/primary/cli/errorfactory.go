package cli

import (
	"errors"
	"strings"
	"time"

	"github.com/Octy-ai/octy-cli/pkg/globals"
	"github.com/briandowns/spinner"
	"github.com/getsentry/sentry-go"
)

// errorFactory: handles outputting error messages with extended help
func errorFactory(errorStr string, spinner *spinner.Spinner) {
	var prefix string
	var msg string = errorStr

	if strings.Contains(errorStr, ":") {
		sections := strings.Split(errorStr, ":")
		prefix = sections[0]
		msg = sections[1]
	}

	switch prefix {
	case "apierror[500]":
		e := newUnknownError(msg, spinner)
		e.quit()
	case "apierror[422]":
		e := newUnprocessableEntityError(msg, spinner)
		e.quit()
	case "apierror[401]":
		e := newUnauthorizedError(msg, spinner)
		e.quit()
	case "apierror[400]":
		e := newBadRequestError(msg, spinner)
		e.quit()
	default:
		logException(errors.New(errorStr))
		e := newUnknownError(globals.ErrUnknownError.Error(), spinner)
		e.quit()
	}
}

//
// Private functions
//

func logException(exception error) {
	sentry.CaptureException(exception)
	sentry.Flush(time.Second * 5)
}
