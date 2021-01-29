# Multitask
[![Build Status](https://travis-ci.org/mhtshhp/multitask.svg)](https://travis-ci.org/mhtshhp/multitask)
[![Go Report Card](https://goreportcard.com/badge/github.com/mhtshhp/multitask)](https://goreportcard.com/report/github.com/mhtshhp/multitask)
[![GoDoc](https://pkg.go.dev/badge/github.com/mhtshhp/multitask?status.svg)](https://pkg.go.dev/github.com/mhtshhp/multitask?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/mhtshhp/multitask/-/badge.svg)](https://sourcegraph.com/github.com/mhtshhp/multitask?badge)
[![Open Source Helpers](https://www.codetriage.com/mhtshhp/multitask/badges/users.svg)](https://www.codetriage.com/mhtshhp/multitask)
[![Release](https://img.shields.io/github/release/mhtshhp/multitask.svg?style=flat-square)](https://github.com/mhtshhp/multitask/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/mhtshhp/multitask)](https://www.tickgit.com/browse?repo=github.com/mhtshhp/multitask)

Multi task for golang


## Features

* Support multitasking efficiently.
* Flexible and easy to use, easy to use

## Usage

See more at [Example](https://github.com/mhtshhp/MultiTask/blob/main/examples/main.go).

```go
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
	fmt.Printf("[ERROR]:%d \n", err)
}

func iterator(assemble []*multitask.Value) {
	for k, v := range assemble {
		fmt.Printf("[M][ID]:%d, [ExcutionTime]:%s, [Error]:%v, [DATA]:%v \n", k, v.ExcutionTime, v.Error, v.Data)
	}
}
```

## License

Released under the [MIT License](https://github.com/mhtshhp/MultiTask/blob/main/LICENSE).
