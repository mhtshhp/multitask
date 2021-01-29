package multitask

import "time"

type Values []Value
type Value struct {
	ID           int
	Data         interface{}
	Error        error
	ExcutionTime time.Duration
}
