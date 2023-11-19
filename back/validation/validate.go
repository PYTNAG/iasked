package validation

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

const (
	InfLowerBound = -1
	InfUpperBound = -1
)

var (
	isUsernameValid = regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString
)

func ValidateStringLength(_string string, min, max int) error {
	length := len(_string)

	if (min != InfLowerBound && length < min) || (max != InfUpperBound && length > max) {
		if min == InfLowerBound {
			msg := "must contain at most %d characters"
			if max == 1 {
				msg, _ = strings.CutSuffix(msg, "s")
			}
			return fmt.Errorf(msg, max)
		}

		if max == InfUpperBound {
			msg := "must contain at least %d characters"
			if min == 1 {
				msg, _ = strings.CutSuffix(msg, "s")
			}
			return fmt.Errorf(msg, min)
		}

		return fmt.Errorf("must contain from %d to %d characters", min, max)
	}

	return nil
}

func ValidateId(id int32) error {
	if id < 1 {
		return fmt.Errorf("must be positive integer")
	}

	return nil
}

func ValidateUsername(username string) error {
	if err := ValidateStringLength(username, 1, InfUpperBound); err != nil {
		return err
	}

	if !isUsernameValid(username) {
		return fmt.Errorf("must contain only eng. letters, digits or underscore")
	}

	return nil
}

func ValidateEmail(email string) error {
	if err := ValidateStringLength(email, 6, InfUpperBound); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("is not a valid email address")
	}

	return nil
}

func ValidatePassword(password string) error {
	if err := ValidateStringLength(password, 6, InfUpperBound); err != nil {
		return err
	}

	return nil
}
