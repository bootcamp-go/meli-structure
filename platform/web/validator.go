package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	// ErrValidatorRequired is the error returned when a field is required
	ErrValidatorRequired = errors.New("validator: field is required")
)

func ValidatorRequiredJSON(r io.Reader, fields ...string) (err error) {
	var m map[string]any

	if err = json.NewDecoder(r).Decode(&m); err != nil {
		return
	}

	for _, field := range fields {
		if _, ok := m[field]; !ok {
			return fmt.Errorf("%w - %s", ErrValidatorRequired, field)
		}
	}

	return
}