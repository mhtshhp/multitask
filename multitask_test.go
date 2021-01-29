package multitask

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

var (
	tasknum int = 200
	start   time.Time
	Printf  = func(format string, v ...interface{}) {}
)

func init() {
	log.SetFlags(log.Llongfile)
	start = time.Now()
	Printf = log.Printf
}

func TestExcute(t *testing.T) {
	fns := make([]MultitaskFunc, tasknum)
	for i := 0; i < tasknum; i++ {
		fns[i] = foreachDemoFunc
	}

	results, err := NewTask(100*time.Millisecond, WithOptionOfQuantityPreExecution(30)).Excute(fns)
	if err != nil {
		t.Fatal("ERROR:", err.Error())
	}
	for k, v := range results {
		var length int
		if v.Data != nil {
			length = len(v.Data.([]byte))
		}
		t.Logf("[M][ID]:%d, [ExcutionTime]:%s, [Error]:%v, [DATA]:%d \n", k, v.ExcutionTime, v.Error, length)
	}
}

func TestTaskManger_Excute(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(4)

	// foreach testing
	go func() {
		defer wg.Done()
		Printf("Single task start")
		for i := 0; i < tasknum; i++ {
			startTime := time.Now()
			data, err := foreachDemoFunc()
			Printf("[S][ID]:%d, [ExcutionTime]:%s, [Error]:%v, [DATA]:%v \n", i, time.Now().Sub(startTime), err, data)
		}
		Printf("SingleTask Process time %s. \n", time.Now().Sub(start))
	}()

	go func() {
		defer wg.Done()
		Printf("Multi task start")
		fns := make([]MultitaskFunc, tasknum)
		for i := 0; i < tasknum; i++ {
			fns[i] = foreachDemoFunc
		}
		fns = append(fns, taskPanicDemo)

		//results, _ := multitask.NewTask(100 * time.Millisecond, multitask.WithOptionOfQuantityPreExecution(1)).Excute(multitask.WithParamsOfMultiFunc(foreachDemoFunc, taskPanicDemo))
		//results, _ := multitask.NewTask(100 * time.Millisecond, multitask.WithOptionOfQuantityPreExecution(1)).Excute(multitask.WithParamsOfFuncMap([]multitask.MultitaskFunc {foreachDemoFunc, taskPanicDemo}))
		//results, _ := multitask.NewTask(100 * time.Millisecond, multitask.WithOptionOfQuantityPreExecution(1)).Excute(multitask.WithParamsOfSingleFunc(foreachDemoFunc))
		results, _ := NewTask(100*time.Millisecond, WithOptionOfQuantityPreExecution(30)).Excute(fns)
		for k, v := range results {
			Printf("[M][ID]:%d, [ExcutionTime]:%s, [Error]:%v, [DATA]:%v \n", k, v.ExcutionTime, v.Error, v.Data)
		}
		Printf("MultiTask Process time %s. \n", time.Now().Sub(start))
	}()

	// request testing
	go func() {
		defer wg.Done()
		Printf("Single task start")
		for i := 0; i < tasknum; i++ {
			startTime := time.Now()
			data, err := request()
			Printf("[S][ID]:%d, [ExcutionTime]:%s, [Error]:%v, [DATA]:%T \n", i, time.Now().Sub(startTime), err, data)
		}
		Printf("SingleTask Process time %s. \n", time.Now().Sub(start))
	}()

	go func() {
		defer wg.Done()
		Printf("Multi task start")
		fns := make([]MultitaskFunc, tasknum)
		for i := 0; i < tasknum; i++ {
			fns[i] = taskRequestDemo
		}
		fns = append(fns, taskPanicDemo)

		//results, _ := multitask.NewTask(100 * time.Millisecond, multitask.WithOptionOfQuantityPreExecution(1)).Excute(multitask.WithParamsOfMultiFunc(foreachDemoFunc, taskPanicDemo))
		//results, _ := multitask.NewTask(100 * time.Millisecond, multitask.WithOptionOfQuantityPreExecution(1)).Excute(multitask.WithParamsOfFuncMap([]multitask.MultitaskFunc {foreachDemoFunc, taskPanicDemo}))
		//results, _ := multitask.NewTask(100 * time.Millisecond, multitask.WithOptionOfQuantityPreExecution(1)).Excute(multitask.WithParamsOfSingleFunc(foreachDemoFunc))
		results, _ := NewTask(100*time.Millisecond, WithOptionOfQuantityPreExecution(30)).Excute(fns)
		for k, v := range results {
			Printf("[M][ID]:%d, [ExcutionTime]:%s, [Error]:%v, [DATA]:%T \n", k, v.ExcutionTime, v.Error, v.Data)
		}
		Printf("MultiTask Process time %s. \n", time.Now().Sub(start))
	}()

	wg.Wait()
}

func BenchmarkSingleExcute(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		foreachDemoFunc()
	}
}

func BenchmarkNormalMultiExcute(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ch := make(chan interface{})
		go func() {
			data, _ := foreachDemoFunc()
			ch <- data
		}()
		select {
		case <-ch:
			//println(data)
		case <-time.After(10 * time.Millisecond):
			println("time out")
		}
	}
}

func BenchmarkMultiExcute(b *testing.B) {

	fns := make([]MultitaskFunc, tasknum)
	fns[0] = foreachDemoFunc

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewTask(10*time.Millisecond, WithOptionOfQuantityPreExecution(tasknum)).Excute(fns)
	}
}

func foreachDemoFunc() (interface{}, error) {

	for i := 0; i < 10000000; i++ {

	}
	return nil, nil
	//return request()
}

func taskPanicDemo() (interface{}, error) {
	panic("BBBBB task b panic BBBBB")
}

func request() ([]byte, error) {
	url := "https://www.google.com/"
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

func taskRequestDemo() (interface{}, error) {
	return request()
}
