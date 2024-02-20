package common

import (
	"errors"
	"sync"
)

// run possibly erroring function `f` concurrently once for each arg
// aggregating errors
func ConcurrentAggregateErrorFn[T any](f func(T) error, args ...T) error {
	errs := []error{}
	errChan := make(chan error)
	success := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(len(args))

	go func() {
		select {
		case err, ok := <-errChan:
			if !ok {
				return
			}
			errs = append(errs, err)
		case <-success:
			return
		}
	}()

	for _, a := range args {
		go func(a T) {
			defer wg.Done()
			err := f(a)
			if err != nil {
				errs = append(errs, err)
			}
		}(a)
	}

	wg.Wait()
	success <- true
	close(errChan)
	close(success)

	return errors.Join(errs...)
}
