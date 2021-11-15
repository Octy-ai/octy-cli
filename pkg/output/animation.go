package output

import (
	"fmt"
	"io"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
)

//
// Spinner
//

type charset = []string

func getCharset() charset {
	// See https://github.com/briandowns/spinner#available-character-sets for
	// list of available charsets
	if runtime.GOOS == "windows" {
		// Less fancy, but uses ASCII characters so works with Windows default
		// console.
		return spinner.CharSets[8]
	}
	return spinner.CharSets[11]
}

const duration = time.Duration(100) * time.Millisecond

// StartNewSpinner starts a new spinner with the given message. If the writer doesn't
// support colors, it simply prints the message.
func StartNewSpinner(msg string, w io.Writer) *spinner.Spinner {
	if !shouldUseColors(w) {
		fmt.Fprintln(w, msg)
		return nil
	}
	s := spinner.New(getCharset(), duration)
	s.Writer = w
	if msg != "" {
		s.Suffix = " " + msg
	}
	s.Start()
	return s
}

// StartSpinner updates an existing spinner's message, and starts it if it was stopped
func StartSpinner(s *spinner.Spinner, msg string, w io.Writer, isStatus bool) {
	if s == nil {
		fmt.Fprintln(w, msg)
		return
	}
	if msg != "" {
		s.Suffix = " " + msg
	}
	if !s.Active() {
		s.Start()
	}
}

// StopSpinner stops a spinner with the given message. If the writer doesn't
// support colors, it simply prints the message.
func StopSpinner(s *spinner.Spinner, msg string, status int, w io.Writer) {
	// if status code, convert to string
	if msg != "" {
		s.FinalMSG = "> " + msg + "\n"
	}
	s.Stop()
}
