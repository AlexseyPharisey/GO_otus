package settings

import (
	"math"
	"os/exec"
	"strconv"
	"strings"
)

func SysMetrics(cpuMetric float64) float64 {
	cores := getCoreCount()
	threads := getActiveThreads()

	normalizedCpuLoad := cpuMetric / (100 * cores)
	normalizedQueue := math.Max(0, threads-cores) / cores
	loadAverage := normalizedCpuLoad + normalizedQueue*0.7

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

func getActiveThreads() float64 {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"(Get-Process | ForEach-Object { $_.Threads.Count } | Measure-Object -Sum).Sum",
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
