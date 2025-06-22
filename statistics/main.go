package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	testMetric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "test27",
		},
		[]string{"label"},
	)
	prometheus.MustRegister(testMetric)

	go func() {
		for {
			testMetric.WithLabelValues("value1").Inc()
			time.Sleep(1 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

//func main() {
//os := runtime.GOOS
//
//keys := []string{"pc", "os", "bios", "cpu", "ram", "disk", "main", "gru", "net"}
//titles := map[string]string{
//	"pc":   "Данные ПК:",
//	"os":   "Система:",
//	"bios": "BIOS:",
//	"cpu":  "Процессор:",
//	"ram":  "Оперативная память:",
//	"disk": "HDD/SSD:",
//	"main": "Материнская плата:",
//	"gru":  "Видеокарта:",
//	"net":  "Сетевые настройки:",
//}
//
//if os == "windows" {
//	result := settings.GetSystemInfo()
//	for _, key := range keys {
//		value, ok := result[key]
//		if !ok {
//			continue
//		}
//		fmt.Printf(titles[key] + "\n")
//
//		switch systemInfoData := value.(type) {
//		case string:
//			fmt.Printf(" - %s\n\n", systemInfoData)
//		case []map[string]any:
//			for _, item := range systemInfoData {
//				for mapKey, mapVal := range item {
//					if mapKey == "Capacity" || mapKey == "Size" || mapKey == "AdapterRAM" {
//						memoryInGb := math.Round(mapVal.(float64) / bytesInGB)
//						fmt.Printf("  - %s: %v GB\n", mapKey, memoryInGb)
//						continue
//					}
//					fmt.Printf("  - %s: %v\n", mapKey, mapVal)
//				}
//				fmt.Println()
//			}
//		case map[string]any:
//			for mapKey, mapVal := range systemInfoData {
//				if mapKey == "Capacity" || mapKey == "Size" || mapKey == "AdapterRAM" {
//					memoryInGb := math.Round(mapVal.(float64) / bytesInGB)
//					fmt.Printf("  - %s: %v GB\n", mapKey, memoryInGb)
//					continue
//				}
//				if key == "os" && mapKey == "Caption" {
//					osName := mapVal.(string)
//					osName = osName[strings.IndexByte(osName, 'W'):]
//					fmt.Printf("  - %s: %v\n", mapKey, osName)
//					continue
//				}
//				fmt.Printf("  - %s: %v\n", mapKey, mapVal)
//			}
//			fmt.Println()
//		default:
//			fmt.Printf("  %v\n\n", systemInfoData)
//		}
//	}
//}
//
//if os == "linux" {
//	fmt.Println(GetAnalyticsLinux())
//}
// Создаем метрику
//}

func prepareSystemData(data map[string]any) {

}
