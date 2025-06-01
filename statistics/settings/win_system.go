package settings

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

var PowerShellCommands = map[string]string{
	"pc":   "Get-WmiObject -Class Win32_SystemEnclosure | Select-Object -ExpandProperty ChassisTypes",
	"net":  "Get-NetIPAddress -AddressFamily IPV4 | Select-Object IPAddress, InterfaceAlias",
	"os":   "Get-CimInstance Win32_OperatingSystem | Select-Object Caption, Manufacturer, Version",
	"bios": "Get-WmiObject -Class Win32_BIOS | Select-Object Name, Version, Manufacturer",
	"cpu":  "Get-WmiObject -Class Win32_Processor | Select-Object Name, Manufacturer, MaxClockSpeed",
	"ram":  "Get-CimInstance -ClassName Win32_PhysicalMemory | Select-Object Capacity, Manufacturer, Speed",
	"disk": "Get-Disk | Select-Object FriendlyName, OperationalStatus, Size",
	"gru":  "Get-CimInstance Win32_VideoController | Select-Object Caption, VideoProcessor, AdapterRAM",
	"main": "Get-CimInstance Win32_BaseBoard | Select-Object Manufacturer, Product",
}

func GetSystemInfo() map[string]any {
	result := make(map[string]any)

	for key, value := range PowerShellCommands {
		commandOutput := sendCommand(value)

		var dataArray []map[string]any
		if err := json.Unmarshal(commandOutput, &dataArray); err == nil {
			result[key] = dataArray
			continue
		}

		var dataMap map[string]any
		if err := json.Unmarshal(commandOutput, &dataMap); err == nil {
			result[key] = dataMap
			continue
		}

		var dataInt int
		if err := json.Unmarshal(commandOutput, &dataInt); err == nil {
			if key == "pc" {
				result[key] = getPcType(dataInt)
				continue
			}
			result[key] = dataInt
			continue
		}

		result[key] = string(commandOutput)
	}

	return result
}

func sendCommand(command string) []byte {
	cmd := exec.Command("powershell", "-Command", command+" | ConvertTo-Json")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("PowerShell Error:", err)
	}

	return output
}

func getPcType(pcType int) string {
	switch pcType {
	case 3:
		return "Desktop"
	case 4:
		return "Low Profile Desktop"
	case 5:
		return "Pizza Box"
	case 6:
		return "Mini Tower"
	case 7:
		return "Tower"
	case 8:
		return "Portable"
	case 9:
		return "Laptop"
	case 10:
		return "Notebook"
	case 11:
		return "Hand Held"
	case 12:
		return "Docking Station"
	case 14:
		return "Sub Notebook"
	case 30:
		return "Tablet"
	default:
		return "Неизвестно"
	}
}
