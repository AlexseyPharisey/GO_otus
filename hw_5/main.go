package main

import (
	"fmt"
)

func main() {
	//SetToChannel()

	goroutinesRunCount := 10
	errorsCount := 2

	goroutines := []Task{
		func() error { return nil },
		func() error { return nil },
		func() error { return ErrErrorsLimitExceeded },
		func() error { return ErrErrorsLimitExceeded },
		func() error { return nil },
		func() error { return ErrErrorsLimitExceeded },
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
		func() error { return ErrErrorsLimitExceeded },
	}

	result := Run(goroutines, goroutinesRunCount, errorsCount)
	fmt.Println(result)
}
