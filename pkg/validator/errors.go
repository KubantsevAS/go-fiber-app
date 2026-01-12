package validator

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatErrors(errors *validate.Errors) string {
	result := ""

	for key, value := range errors.Errors {
		result += fmt.Sprintf("%s: %s\n", key, strings.Join(value, ", "))
	}

	return result
}
