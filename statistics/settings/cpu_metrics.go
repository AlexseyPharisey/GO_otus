package settings

import (
	"os/exec"
	"strconv"
	"strings"
)

func CpuMetrics(interval int) float64 {
	count := 0
	sum := 0

	for i := 0; i < 2; i++ {
		cmd := exec.Command(
			"powershell",
			"-Command",
			"Get-WmiObject Win32_Processor | Select-Object -ExpandProperty LoadPercentage",
		)

		output, err := cmd.Output()
		if err != nil {
			continue
		}

		loadStr := strings.TrimSpace(string(output))
		load, err := strconv.Atoi(loadStr)
		if err != nil {
			continue
		}

		sum += load
		count++
	}
	avg := float64(sum) / float64(interval)
	return avg / 100
}
