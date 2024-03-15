package tools

import "errors"

var (
	ErrFieldNotFound = errors.New("field not found")
)

func CheckFieldExistance(fields map[string]any, requiredFields ...string) (err error) {
	for _, field := range requiredFields {
		if _, ok := fields[field]; !ok {
			return ErrFieldNotFound
		}
	}
	return nil
}
