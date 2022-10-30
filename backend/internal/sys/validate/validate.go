package validate

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Checks a struct for validation errors and returns any errors the occur. This
// wraps the validate.Struct() function and provides some error wrapping. When
// a validator.ValidationErrors is returned, it is wrapped transformed into a
// FieldErrors array and returned.
func Check(val any) error {
	err := validate.Struct(val)

	if err != nil {
		verrors, ok := err.(validator.ValidationErrors)
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
