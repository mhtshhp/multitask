package main

import (
	"errors"
	"fmt"
	"github.com/mhtshhp/multitask"
	"time"
)

func dosamething0() (interface{}, error) {
	// dosamething
	time.Sleep(time.Second)
	return nil, nil
}

func dosamething1() (interface{}, error) {
	// dosamething
	time.Sleep(2 * time.Second)
	return nil, nil
}

func dosamething2() (interface{}, error) {
	// dosamething
	panic(errors.New("unknown error"))
	return nil, nil
}

func main() {

	var (
		results    []*multitask.Value
		err        error
		timeout    = 100 * time.Millisecond
		concurrent = multitask.WithOptionOfQuantityPreExecution(2)
	)

	// Different calling methods are supported

	// single task usage
	// 1
	results, err = multitask.NewTask(timeout, concurrent).Excute(multitask.WithParamsOfSingleFunc(dosamething0))
	if err != nil {
		handleError(err)
	}
	iterator(results)

	// multi task usage
	// 2
	results, err = multitask.NewTask(timeout, concurrent).Excute(multitask.WithParamsOfMultiFunc(dosamething0, dosamething1, dosamething2))
	if err != nil {
		handleError(err)
	}
	iterator(results)

	// 3
	results, err = multitask.NewTask(timeout, concurrent).Excute(multitask.WithParamsOfFuncMap([]multitask.MultitaskFunc{dosamething0, dosamething1, dosamething2}))
	if err != nil {
		handleError(err)
	}
	iterator(results)

	// 4
	fns := make([]multitask.MultitaskFunc, 3)
	fns[0] = dosamething0
	fns[1] = dosamething1
	fns[2] = dosamething2
	results, err = multitask.NewTask(timeout, concurrent).Excute(fns)
	if err != nil {
		handleError(err)
	}
	iterator(results)

}

func handleError(err error) {
	fmt.Printf("[ERROR]:%v \n", err)
}

func iterator(assemble []*multitask.Value) {
	for k, v := range assemble {
		fmt.Printf("[M][ID]:%d, [ExcutionTime]:%s, [Error]:%v, [DATA]:%v \n", k, v.ExcutionTime, v.Error, v.Data)
	}
}
