package settings

import (
	"os/exec"
	"strconv"
	"strings"
)

func RamMetrics(interval int) float64 {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"[math]::Round((($mem = Get-CimInstance Win32_OperatingSystem) |"+
			" % { ($_.TotalVisibleMemorySize - $_.FreePhysicalMemory) / $_.TotalVisibleMemorySize }) * 100, 2)",
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
