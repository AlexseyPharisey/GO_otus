package main

import (
	"GO_otus/statistics/settings"
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"
)

const N = 5
const M = 5

func main() {
	diskNames := settings.GetDiskName()
	diskStats := make(map[string]map[string][]float64)
	result := make(map[string]map[string]float64)
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

	go func() {
		for {
			cpu := settings.CpuMetrics()
			mu.Lock()
			cpuStats = append(cpuStats, cpu)
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			mu.Lock()
			for _, disk := range diskNames {
				speed := settings.DiskSpeedMetrics(disk)
				mem := settings.DiskMemoryMetrics(disk)

				if speed != nil {
					diskStats[disk]["KBps"] = append(diskStats[disk]["KBps"], speed["KBps"])
					diskStats[disk]["TPS"] = append(diskStats[disk]["TPS"], speed["TPS"])
				}
				if mem != nil {
					diskStats[disk]["UsedPercent"] = append(
						diskStats[disk]["UsedPercent"],
						mem["UsedPercent"])
					diskStats[disk]["UsedInodePercent"] = append(
						diskStats[disk]["UsedInodePercent"],
						mem["UsedInodePercent"])
				}
			}
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		outputTicker := time.NewTicker(N * time.Second)
		defer outputTicker.Stop()

		initialized := false
		for range outputTicker.C {
			mu.Lock()

			if !initialized {
				cpuReady := len(cpuStats) >= M
				diskReady := len(diskNames) > 0

				for _, disk := range diskNames {
					if len(diskStats[disk]["KBps"]) < M ||
						len(diskStats[disk]["TPS"]) < M ||
						len(diskStats[disk]["UsedPercent"]) < M ||
						len(diskStats[disk]["UsedInodePercent"]) < M {
						diskReady = false
						break
					}
				}

				if !(cpuReady && diskReady) {
					mu.Unlock()
					continue
				}
				initialized = true
			}

			result["cpu"] = map[string]float64{
				"cpu_load": prepareAvg(cpuStats[len(cpuStats)-M:]),
			}

			result["system"] = map[string]float64{
				"load_average": settings.SysMetrics(result["cpu"]["cpu_load"]),
			}

			for _, disk := range diskNames {
				result[disk] = map[string]float64{
					"KBps": prepareAvg(diskStats[disk]["KBps"][len(diskStats[disk]["KBps"])-M:]),
					"TPS":  prepareAvg(diskStats[disk]["TPS"][len(diskStats[disk]["TPS"])-M:]),
					"UsedPercent": prepareAvg(
						diskStats[disk]["UsedPercent"][len(diskStats[disk]["UsedPercent"])-M:]) * 100,
					"UsedInodePercent": prepareAvg(
						diskStats[disk]["UsedInodePercent"][len(diskStats[disk]["UsedInodePercent"])-M:]) * 100,
				}
			}

			jsonData, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				fmt.Println("JSON error:", err)
			} else {
				fmt.Println(string(jsonData))
			}
			mu.Unlock()
		}
	}()

	select {}
}

func prepareAvg(arr []float64) float64 {
	if len(arr) < M {
		return 0
	}
	var sum float64
	for i := 0; i < M; i++ {
		sum += arr[i]
	}
	avg := sum / float64(M)

	return math.Round((avg/100)*100) / 100
}
