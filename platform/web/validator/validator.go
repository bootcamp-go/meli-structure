package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	// ErrRequired is the error returned when a field is required
	ErrRequired = errors.New("validator: field is required")
)

func RequiredJSON(r io.Reader, fields ...string) (err error) {
	var m map[string]any

	if err = json.NewDecoder(r).Decode(&m); err != nil {
		return
	}

	for _, field := range fields {
		if _, ok := m[field]; !ok {
			return fmt.Errorf("%w - %s", ErrRequired, field)
		}
	}

	return
}