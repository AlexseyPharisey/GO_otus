package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors Limit Exceeded")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	if len(tasks) == 0 {
		return nil
	}

	var errorCount int64 = 0
	stopCh := make(chan struct{})
	taskCh := make(chan Task)
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-stopCh:
					fmt.Printf("Горутина %d остановлена\n", id)
					return
				case task, ok := <-taskCh:
					if !ok {
						fmt.Printf("Горутина %d: задачи закончились\n", id)
						return
					}

					err := task()
					if err != nil {
						newCount := atomic.AddInt64(&errorCount, 1)
						fmt.Printf("Горутина %d: ошибка (всего %d)\n", id, newCount)
						if newCount >= int64(m) {
							close(stopCh)
							return
						}
					}
				}
			}
		}(i)
	}

	go func() {
		defer close(taskCh)
		for _, task := range tasks {
			select {
			case <-stopCh:
				fmt.Println("Остановка отправки задач")
				return
			case taskCh <- task:
			}
		}
	}()

	wg.Wait()

	if atomic.LoadInt64(&errorCount) >= int64(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
