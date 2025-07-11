package main

import (
	"GO_otus/statistics/goroutines"
	"GO_otus/statistics/settings"
	"GO_otus/statistics/storage"
	"fmt"
	"sync"
	"time"
)

const N = 5
const M = 15

func main() {
	getMetrics()
	select {}
}

func getMetrics() {
	diskNames := settings.GetDiskName()
	diskStats := make(map[string]map[string][]float64)
	var cpuStats []float64
	var mu sync.Mutex

	for _, d := range diskNames {
		diskStats[d] = map[string][]float64{
			"KBps":             {},
			"TPS":              {},
			"UsedPercent":      {},
			"UsedInodePercent": {},
		}
	}

	report := &storage.Report{}

	goroutines.CpuRoutine(&mu, &cpuStats)
	goroutines.DiskRoutine(&mu, diskNames, diskStats)
	goroutines.ResultRoutine(&mu, diskNames, &cpuStats, diskStats,
		make(map[string]map[string]float64), N, M, report)

	ticker := time.NewTicker(time.Duration(N) * time.Second)

	for range ticker.C {
		if data := report.Get(); data != nil {
			fmt.Println(string(data))
		}
	}
}
