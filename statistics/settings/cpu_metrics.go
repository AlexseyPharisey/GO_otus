package settings

import (
	"os/exec"
	"strconv"
	"strings"
)

func CpuMetrics() float64 {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"Get-WmiObject Win32_Processor | Select-Object -ExpandProperty LoadPercentage",
	)

	output, err := cmd.Output()
	if err != nil {
		return 0
	}

	loadStr := strings.TrimSpace(string(output))
	load, err := strconv.Atoi(loadStr)
	if err != nil {
		return 0
	}

	return float64(load)
}
