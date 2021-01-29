package multitask

import (
	"errors"
)

var (
	// Invalid
	ErrInvalidExpiryDuration       = errors.New("[MultiTask] invalid Expiry Duration")
	ErrInvalidQuantityPreExecution = errors.New("[MultiTask] invalid Quantity Pre Execution")

	// Excute
	ErrExcuteTimeout = errors.New("[MultiTask] excute timeout")
	ErrExcutePanic   = errors.New("[MultiTask] excute panic: ")
	ErrExcuteUnknown = errors.New("[MultiTask] excute unknown error")
)

func mergeErrors(errs ...error) error {
	newErrs := ""
	for _, err := range errs {
		newErrs += "" + err.Error() + " "
	}
	return errors.New(newErrs)
}
