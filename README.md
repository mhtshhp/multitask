# Multitask
[![Build Status](https://travis-ci.org/mhtshhp/multitask.svg)](https://travis-ci.org/mhtshhp/multitask)
[![Go Report Card](https://goreportcard.com/badge/github.com/mhtshhp/multitask)](https://goreportcard.com/report/github.com/mhtshhp/multitask)
[![GoDoc](https://pkg.go.dev/badge/github.com/mhtshhp/multitask?status.svg)](https://pkg.go.dev/github.com/mhtshhp/multitask?tab=doc)
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
		concurrent = multitask.WithOptionOfQuantityPreExecution(