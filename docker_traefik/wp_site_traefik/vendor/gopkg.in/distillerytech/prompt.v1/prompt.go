package prompt

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var out io.Writer = os.Stdout
var in io.Reader = os.Stdin

// AskString prompts for a string using provided prompt.
func AskString(prompt string) string {
	var result string

	fmt.Fprint(out, prompt+" ")
	fmt.Fscanln(in, &result)

	return result
}

// AskStringRequired prompts for a required string. Prompt will repeat until
// a nonempty input is entered.
func AskStringRequired(prompt string) string {
	result := AskString(prompt)

	if strings.TrimSpace(result) == "" {
		return AskStringRequired(prompt)
	}

	return result
}

// AskStringLimit prompts for a string limited to a list of valid inputs.
// If no valid inputs are provided, any input is valid.
func AskStringLimit(prompt string, validInputs ...string) string {
	result := AskString(prompt)

	if !validate(result, validInputs...) {
		return AskStringLimit(prompt, validInputs...)
	}

	return result
}

// AskInteger prompts for an integer using provided prompt. Prompt
// will repeat until a valid integer is entered.
func AskInteger(prompt string) int {
	return integer(prompt, false, 0)
}

// AskIntegerDefault prompts for an integer using provided prompt. The
// default value will be returned if invalid integer is entered.
func AskIntegerDefault(prompt string, defaultVal int) int {
	return integer(prompt, true, defaultVal)
}

// Confirm prompts for bool. Prompt will repeat unilt a valid input is entered
// The following are considered valid inputs:
// y, Y, yes, Yes, n, N, no, No.
func Confirm(prompt string) bool {
	result := AskString(prompt)

	switch result {
	case "y", "Y", "yes", "Yes":
		return true
	case "n", "N", "no", "No":
		return false
	default:
		return Confirm(prompt)
	}

}

func integer(prompt string, useDefault bool, defaultVal int) int {
	result := AskString(prompt)

	i, err := strconv.Atoi(result)

	switch {
	case err == nil:
		return i
	case useDefault:
		return defaultVal
	default:
		return integer(prompt, useDefault, defaultVal)
	}
}

func validate(input string, validInputs ...string) bool {

	if len(validInputs) > 0 {
		for _, validInput := range validInputs {
			if input == validInput {
				return true
			}
		}

		return false
	}

	return true
}
