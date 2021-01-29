package multitask

import (
	"errors"
	"fmt"
	"time"
)

func handleMultiTask(task_id int, fn func() (interface{}, error), timeout time.Duration, ch chan interface{}) {
	startTime := time.Now()
	ch_run := make(chan interface{})
	go handleTask(task_id, ch_run, fn)
	select {
	case re := <-ch_run:
		endTime := time.Now()
		re.(*Value).ExcutionTime = endTime.Sub(startTime)
		ch <- re

	case <-time.After(timeout):
		endTime := time.Now()
		resulter := &Value{
			ID:           task_id,
			Data:         nil,
			Error:        ErrExcuteTimeout,
			ExcutionTime: endTime.Sub(startTime),
		}
		ch <- resulter

	}
}

func handleTask(task_id int, ch chan interface{}, fn func() (interface{}, error)) {
	defer func() {
		if err := recover(); err != nil {
			resulter := &Value{
				ID:    task_id,
				Data:  nil,
				Error: mergeErrors(ErrExcutePanic, errors.New(fmt.Sprintf("%s", err))),
			}
			ch <- resulter
		}
	}()

	data, err := fn()

	resulter := &Value{
		ID:    task_id,
		Data:  data,
		Error: err,
	}
	ch <- resulter

	return
}

type taskManger struct {
	ExpiryDuration time.Duration
	Option         taskOption
}

type MultitaskFuncChains []MultitaskFunc

// multi task func
type MultitaskFunc func() (interface{}, error) // ([]*multitask.Resulter, error)

// new task public instance
func NewTask(expiryDuration time.Duration, opts ...TaskOption) *taskManger {
	jobManager := &taskManger{
		ExpiryDuration: expiryDuration,
		Option:         defaultOptions(),
	}
	for _, opt := range opts {
		opt.apply(&jobManager.Option)
	}
	return jobManager
}

// excute all kind of function type
// map, collection and single funciton
func (t *taskManger) Excute(fnCollection []MultitaskFunc) ([]*Value, error) {
	return t.excute(fnCollection)
}

func (t *taskManger) excute(fnCollection []MultitaskFunc) ([]*Value, error) {
	// Options
	if t.ExpiryDuration <= 0 {
		return nil, ErrInvalidExpiryDuration
	}
	if t.Option.QuantityPreExecution <= 0 {
		t.Option.QuantityPreExecution = len(fnCollection)
	}
	timeout := t.ExpiryDuration
	maxTaskCount := t.Option.QuantityPreExecution
	chLimit := make(chan bool, maxTaskCount)
	chs := make([]chan interface{}, len(fnCollection))
	limitFunc := func(chLimit chan bool, ch chan interface{}, task_id int, fn func() (interface{}, error), timeout time.Duration) {
		defer func() {
			if err := recover(); err != nil {
				<-chLimit
			}
		}()
		handleMultiTask(task_id, fn, timeout, ch)
		<-chLimit
	}
	for i, fn := range fnCollection {
		chs[i] = make(chan interface{}, 1)
		chLimit <- true
		go limitFunc(chLimit, chs[i], i, fn, timeout)
	}
	values := make([]*Value, len(fnCollection))
	for i, ch := range chs {
		resulter := <-ch
		values[i] = resulter.(*Value)
	}
	return values, nil
}
