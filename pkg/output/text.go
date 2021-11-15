package output

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

//
// Public functions
//

var BPrint = color.New(color.Bold).PrintlnFunc()

var FPrint = color.New(color.Bold, color.FgHiRed).PrintlnFunc()

var SPrint = color.New(color.Bold, color.FgGreen).PrintlnFunc()

// Linkify returns an ANSI escape sequence with an hyperlink, if the writer
// supports colors.
func Linkify(text, url string, w io.Writer) string {
	if !shouldUseColors(w) {
		return text
	}

	// See https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
	// for more information about this escape sequence.
	return fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", url, text)
}

// ForceColors forces the use of colors and other ANSI sequences.
var ForceColors = false

// DisableColors disables all colors and other ANSI sequences.
var DisableColors = false

// EnvironmentOverrideColors overs coloring based on `CLICOLOR` and
// `CLICOLOR_FORCE`. Cf. https://bixense.com/clicolors/
var EnvironmentOverrideColors = true

//
// Private functions
//

// ColorizeStatus prints a colorized number for HTTP status code
func ColorizeStatus(status int) {

	switch {
	case status >= 500:
		color.Red(strconv.Itoa(status))
	case status >= 300:
		color.Yellow(strconv.Itoa(status))
	default:
		color.Green(strconv.Itoa(status))
	}
}

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}

func shouldUseColors(w io.Writer) bool {
	useColors := ForceColors || checkIfTerminal(w)

	if EnvironmentOverrideColors {
		force, ok := os.LookupEnv("CLICOLOR_FORCE")

		switch {
		case ok && force != "0":
			useColors = true
		case ok && force == "0":
			useColors = false
		case os.Getenv("CLICOLOR") == "0":
			useColors = false
		}
	}

	return useColors && !DisableColors
}
