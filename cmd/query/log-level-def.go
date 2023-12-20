package query

import (
	"errors"
	"fmt"
)

type logLevelEnum string

const (
	errorLogLevel logLevelEnum = "error"
	infoLogLevel  logLevelEnum = "info"
)

// String is used both by fmt.Print and by Cobra in help text
func (e *logLevelEnum) String() string {
	return string(*e)
}

// Set must have pointer receiver, so it doesn't change the value of a copy
func (e *logLevelEnum) Set(v string) error {
	switch v {
	case "error", "info":
		*e = logLevelEnum(v)
		return nil
	default:
		return errors.New(
			fmt.Sprintf(`must be one of "%s" or "%s"`, errorLogLevel, infoLogLevel),
		)
	}
}

// Type is only used in help text
func (e *logLevelEnum) Type() string {
	return "logLevelEnum"
}
