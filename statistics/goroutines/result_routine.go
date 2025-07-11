package goroutines

import (
	"GO_otus/statistics/settings"
	"encoding/json"
	"math"
	"sync"
	"time"
)

func ResultRoutine(mu *sync.Mutex,
	diskNames []string,
	cpuStats *[]float64,
	diskStats map[string]map[string][]float64,
	result map[string]map[string]float64,
	N, M int,
	store interface{ Set([]byte) },
) {
	go func() {
		ticker := time.NewTicker(time.Duration(N) * time.Second)
		defer ticker.Stop()

		initialized := false

		for range ticker.C {
			mu.Lock()

			if !initialized {
				cpuReady := len(*cpuStats) >= M
				diskReady := true
				for _, d := range diskNames {
					if len(diskStats[d]["KBps"]) < M ||
						len(diskStats[d]["TPS"]) < M ||
						len(diskStats[d]["UsedPercent"]) < M ||
						len(diskStats[d]["UsedInodePercent"]) < M {
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
				"cpu_load_percent": round0(prepareAvg((*cpuStats)[len(*cpuStats)-M:], M)),
			}
			result["system"] = map[string]float64{
				"load_average": settings.SysMetrics(result["cpu"]["cpu_load_percent"]),
			}
			for _, d := range diskNames {
				ds := diskStats[d]
				result[d] = map[string]float64{
					"KBps":             round2(prepareAvg(ds["KBps"][len(ds["KBps"])-M:], M)),
					"TPS":              round2(prepareAvg(ds["TPS"][len(ds["TPS"])-M:], M)),
					"UsedPercent":      round0(prepareAvg(ds["UsedPercent"][len(ds["UsedPercent"])-M:], M)),
					"UsedInodePercent": round0(prepareAvg(ds["UsedInodePercent"][len(ds["UsedInodePercent"])-M:], M)),
				}
			}

			if js, err := json.Marshal(result); err == nil {
				store.Set(js)
			}

			mu.Unlock()
		}
	}()
}

func prepareAvg(arr []float64, M int) float64 {
	if len(arr) == 0 {
		return 0
	}
	var sum float64
	for _, v := range arr {
		sum += v
	}
	avg := sum / float64(M)
	return avg
}

func round2(x float64) float64 {
	return math.Round(x*100) / 100
}
func round0(x float64) float64 {
	return math.Round(x)
}
