// Package validate provides a wrapper around the go-playground/validator package
package validate

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() { // nolint
	validate = validator.New()

	err := validate.RegisterValidation("shoutrrr", func(fl validator.FieldLevel) bool {
		prefixes := [...]string{
			"bark://",
			"discord://",
			"smtp://",
			"gotify://",
			"googlechat://",
			"ifttt://",
			"join://",
			"mattermost://",
			"matrix://",
			"ntfy://",
			"opsgenie://",
			"pushbullet://",
			"pushover://",
			"rocketchat://",
			"slack://",
			"teams://",
			"telegram://",
			"zulip://",
			"generic://",
			"generic+",
		}

		str := fl.Field().String()
		if str == "" {
			return false
		}

		for _, prefix := range prefixes {
			if strings.HasPrefix(str, prefix) {
				return true
			}
		}

		return false
	})

	if err != nil {
		panic(err)
	}
}

// Check a struct for validation errors and returns any errors the occur. This
// wraps the validate.Struct() function and provides some error wrapping. When
// a validator.ValidationErrors is returned, it is wrapped transformed into a
// FieldErrors array and returned.
func Check(val any) error {
	err := validate.Struct(val)
	if err != nil {
		verrors, ok := err.(validator.ValidationErrors) // nolint - we know it's a validator.ValidationErrors
		if !ok {
			return err
		}

		fields := make(FieldErrors, 0, len(verrors))
		for _, verr := range verrors {
			field := FieldError{
				Field: verr.Field(),
				Error: verr.Error(),
			}

			fields = append(fields, field)
		}
		return fields
	}

	return nil
}
