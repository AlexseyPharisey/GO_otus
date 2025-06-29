package settings

import (
	"os/exec"
	"strconv"
	"strings"
)

func NetMetrics(interval int) float64 {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"Get-Counter \"\\Network Interface(*)\\Bytes Total/sec\""+
			" | % { $_.CounterSamples | % { \"$($_.Path.Split('\\\\')[-2]): \""+
			" + [math]::Round(($_.CookedValue / (100 * 125000)) * 100, 2) + \" %\" } }\n",
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
