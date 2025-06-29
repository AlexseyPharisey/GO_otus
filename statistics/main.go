package main

import "GO_otus/statistics/settings"

const bytesInGB = 1024 * 1024 * 1024

func main() {
	//N := 5
	M := 15

	var cpuResults []float64
	//var sysResults []float64
	var ramResults []float64
	var diskResults []float64
	var netResults []float64
	for i := 0; i < 4; i++ {
		cpuMetrics := settings.CpuMetrics(M)
		cpuResults = append(cpuResults, cpuMetrics)

		//sysMetrics := settings.SysMetrics(cpuMetrics)
		//sysResults = append(sysResults, sysMetrics)

		ramMetrics := settings.RamMetrics(M)
		ramResults = append(ramResults, ramMetrics)

		diskMetrics := settings.DiskMetrics(M)
		diskResults = append(diskResults, diskMetrics)

		netMetrics := settings.NetMetrics(M)
		netResults = append(netResults, netMetrics)
	}

	println("CPU Results: \n", cpuResults)
	//println("System Results: ", sysResults)
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
