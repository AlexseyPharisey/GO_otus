package main

import (
	"GO_otus/statistics/settings"
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"
)

//func main() {
//	counter := 0
//	var arr []int
//
//	for {
//		arr = append(arr, counter)
//		counter++
//
//		if len(arr) > 10 {
//			arr = arr[len(arr)-5:]
//		}
//
//		fmt.Println(arr)
//		time.Sleep(2 * time.Second)
//	}
//}

func main() {
	diskNames := []string{"C", "D", "E"}
	diskStats := make(map[string]map[string][]float64)
	result := make(map[string]map[string]float64)
	var cpuStats []float64
	var mu sync.Mutex

	// Инициализация структур для хранения данных
	for _, d := range diskNames {
		diskStats[d] = map[string][]float64{
			"KBps":             {},
			"TPS":              {},
			"UsedPercent":      {},
			"UsedInodePercent": {},
		}
	}

	// Горутина для сбора CPU метрик
	go func() {
		for {
			cpu := settings.CpuMetrics()
			mu.Lock()
			cpuStats = append(cpuStats, cpu)
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	// Горутина для сбора Disk метрик
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
		ticker := time.NewTicker(1 * time.Second) // Проверяем каждую секунду
		defer ticker.Stop()

		for range ticker.C {
			mu.Lock()

			// Проверяем наличие достаточного количества данных (ваша оригинальная логика)
			cpuReady := len(cpuStats) == 15 // 15-1 как у вас было
			diskReady := false
			for _, disk := range diskNames {
				for _, values := range diskStats[disk] {
					if len(values) == 15 { // 15-2 как у вас было
						diskReady = true
						break
					}
				}
				if diskReady {
					break
				}
			}

			if cpuReady && diskReady {
				result["cpu"] = map[string]float64{
					"cpu_load": prepareAvg(cpuStats),
				}
				result["system"] = map[string]float64{
					"load_average": settings.SysMetrics(result["cpu"]["cpu_load"]),
				}
				for _, disk := range diskNames {
					result[disk] = map[string]float64{
						"KBps":             prepareAvg(diskStats[disk]["KBps"]),
						"TPS":              prepareAvg(diskStats[disk]["TPS"]),
						"UsedPercent":      prepareAvg(diskStats[disk]["UsedPercent"]) * 100,
						"UsedInodePercent": prepareAvg(diskStats[disk]["UsedInodePercent"]) * 100,
					}
				}
				jsonData, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					fmt.Println("JSON error:", err)
				} else {
					fmt.Println(string(jsonData))
				}

				cpuStats = cpuStats[5:]
				for _, disk := range diskNames {
					for key := range diskStats[disk] {
						diskStats[disk][key] = diskStats[disk][key][5:]
					}
				}
			}

			mu.Unlock()
		}
	}()

	// Блокируем main(), чтобы программа не завершилась
	select {}
}

func prepareAvg(arr []float64) float64 {
	if len(arr) < 15 {
		return 0
	}
	var sum float64
	for i := 0; i < 15; i++ {
		sum += arr[i]
	}
	avg := sum / float64(15)

	return math.Round((avg/100)*100) / 100
}
