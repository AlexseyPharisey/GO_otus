package settings

import (
	"os/exec"
	"strconv"
	"strings"
)

func DiskMetrics(interval int) float64 {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"Get-CimInstance Win32_LogicalDisk -Filter \"DriveType=3\""+
			" | % { \"$($_.DeviceID): \" + [math]::Round((($_.Size - $_.FreeSpace) / $_.Size) * 100, 2) + \" %\" }",
	)

	output, err := cmd.Output()
	if err != nil {
		return 0
	}

	loadStr := strings.TrimSpace(string(output))
	loadStr = strings.ReplaceAll(loadStr, ",", ".") // заменяем запятую на точку
	load, err := strconv.ParseFloat(loadStr, 64)
	if err != nil {
		return 0
	}

	return float64(load)
}
