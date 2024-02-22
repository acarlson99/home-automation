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

	wgA := sync.WaitGroup{}
	wgA.Add(1)
	go func() {
		defer wgA.Done()
		for err := range errChan {
			errs = append(errs, err)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(len(args))

	for _, a := range args {
		go func(a T) {
			defer wg.Done()
			err := f(a)
			if err != nil {
				errChan <- err
			}
		}(a)
	}

	wg.Wait()
	close(errChan)
	wgA.Wait()

	return errors.Join(errs...)
}

func AggregateErrorFn[T any](f func(T) error, args ...T) error {
	errs := []error{}

	for _, a := range args {
		func(a T) {
			err := f(a)
			if err != nil {
				errs = append(errs, err)
			}
		}(a)
	}

	return errors.Join(errs...)
}
