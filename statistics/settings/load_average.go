package settings

import (
	"math"
	"os/exec"
	"strconv"
	"strings"
)

func SysMetrics(cpuMetric float64) float64 {
	cores := getCoreCount()
	queue := getQueue()
	loadAverage := (cpuMetric / 100) * (1 + queue/cores)

	return math.Round(loadAverage*100) / 100
}

func getCoreCount() float64 {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"(Get-WmiObject Win32_Processor).NumberOfLogicalProcessors",
	)

	output, err := cmd.Output()
	if err != nil {
		return 0
	}

	outputStr := strings.TrimSpace(string(output))
	outputInt, err := strconv.Atoi(outputStr)
	if err != nil {
		return 0
	}

	return float64(outputInt)
}

func getQueue() float64 {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"(Get-CimInstance Win32_PerfFormattedData_PerfOS_System).ProcessorQueueLength",
	)

	output, err := cmd.Output()
	if err != nil {
		return 0
	}

	outputStr := strings.TrimSpace(string(output))
	outputInt, err := strconv.Atoi(outputStr)
	if err != nil {
		return 0
	}

	return float64(outputInt)
}
