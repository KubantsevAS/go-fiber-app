package validator

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatErrors(errors *validate.Errors) string {
	result := ""

	for _, value := range errors.Errors {
		result += fmt.Sprintf("%s\n", strings.Join(value, ", "))
	}

	return result
}
