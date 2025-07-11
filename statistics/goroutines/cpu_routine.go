package goroutines

import (
	"GO_otus/statistics/settings"
	"sync"
	"time"
)

func CpuRoutine(mu *sync.Mutex, cpuStats *[]float64) {
	go func() {
		for {
			cpu := settings.CpuMetrics()
			mu.Lock()
			*cpuStats = append(*cpuStats, cpu)
			mu.Unlock()

			time.Sleep(1 * time.Second)
		}
	}()
}
