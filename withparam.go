package multitask

func WithParamsOfMultiFunc(fnCollection ...MultitaskFunc) []MultitaskFunc {
	return fnCollection
}

func WithParamsOfFuncMap(fnCollection []MultitaskFunc) []MultitaskFunc {
	return fnCollection
}

func WithParamsOfSingleFunc(fnCollection MultitaskFunc) []MultitaskFunc {
	return []MultitaskFunc{fnCollection}
}
