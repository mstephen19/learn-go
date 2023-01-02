package lib

type ReduceHandler[Acc comparable, Type comparable] func(acc *Acc, item Type, index int)

// Run a handler for every item in a slice while reducing down to a new single value.
func Reduce[Acc comparable, Type comparable](slice []Type, acc Acc, handler ReduceHandler[Acc, Type]) *Acc {
	for index, item := range slice {
		handler(&acc, item, index)
	}

	return &acc;
}