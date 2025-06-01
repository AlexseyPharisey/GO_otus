package main

import (
	"GO_otus/statistics/settings"
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS

	keys := []string{
		"pc_settings",
		"os_settings",
		"bios_settings",
		"cpu_settings",
		"ram_settings",
		"disk_settings",
		"main_settings",
		"gru_settings",
		"net_settings",
	}

	if os == "windows" {
		result := settings.SystemInfo()
		for _, key := range keys {
			value, ok := result[key]
			if !ok {
				continue
			}
			switch key {
			case "pc_settings":
				fmt.Println("Данные ПК:")
			case "os_settings":
				fmt.Println("Система:")
			case "bios_settings":
				fmt.Println("BIOS:")
			case "cpu_settings":
				fmt.Println("Процессор:")
			case "ram_settings":
				fmt.Println("Оперативная память:")
			case "disk_settings":
				fmt.Println("HDD/SSD:")
			case "main_settings":
				fmt.Println("Материнская плата:")
			case "gru_settings":
				fmt.Println("Видеокарта:")
			case "net_settings":
				fmt.Println("Сетевые настройки:")
			default:
				fmt.Printf("%s:\n", key)
			}

			switch v := value.(type) {
			case string:
				fmt.Printf(" - %s\n\n", v)
			case []map[string]any:
				for _, item := range v {
					for k, val := range item {
						fmt.Printf("  - %s: %v\n", k, val)
					}
					fmt.Println()
				}
			case map[string]any:
				for k, val := range v {
					fmt.Printf("  - %s: %v\n", k, val)
				}
				fmt.Println()
			default:
				fmt.Printf("  %v\n\n", v)
			}
		}
	}

	if os == "linux" {
		fmt.Println(GetAnalyticsLinux())
	}
}
