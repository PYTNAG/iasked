package validation

import "fmt"

const (
	InfLowerBound = -1
	InfUpperBound = -1
)

func validateStringLength(_string string, min, max int) error {
	length := len(_string)

	if (min != InfLowerBound && length < min) || (max != InfUpperBound && length > max) {
		if min == InfLowerBound {
			return fmt.Errorf("must contain at most %d characters", max)
		}

		if max == InfUpperBound {
			return fmt.Errorf("must contain at least %d characters", min)
		}

		return fmt.Errorf("must contain from %d to %d characters", min, max)
	}

	return nil
}

func validateId(id int32) error {
	if id < 1 {
		return fmt.Errorf("must be positive integer")
	}

	return nil
}
