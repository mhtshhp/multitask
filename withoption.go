package multitask

//task option
type taskOption struct {
	QuantityPreExecution int
}

type TaskOption interface {
	apply(*taskOption)
}

type funcOption struct {
	f func(*taskOption)
}

func (fdo *funcOption) apply(do *taskOption) {
	fdo.f(do)
}

func newFuncOption(f func(*taskOption)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithOptionOfQuantityPreExecution(num int) TaskOption {
	return newFuncOption(func(o *taskOption) {
		o.QuantityPreExecution = num
	})
}

//default options
func defaultOptions() taskOption {
	return taskOption{
		QuantityPreExecution: 0,
	}
}
