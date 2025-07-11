package goroutines

import (
	"GO_otus/statistics/settings"
	"sync"
	"time"
)

func DiskRoutine(
	mu *sync.Mutex,
	diskNames []string,
	diskStats map[string]map[string][]float64,
) {
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
}
