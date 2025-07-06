package settings

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func DiskSpeedMetrics(diskName string) map[string]float64 {
	psScript := fmt.Sprintf(`
		$disk="%s:";
		$d=Get-CimInstance Win32_PerfFormattedData_PerfDisk_LogicalDisk | Where-Object { $_.Name -eq $disk };
		if ($d) {
			@{
				TPS=[math]::Round($d.DiskTransfersPersec,2);
				KBps=[math]::Round(($d.DiskReadBytesPerSec + $d.DiskWriteBytesPerSec)/1024,2)
			} | ConvertTo-Json -Depth 2
		}
	`, diskName)

	cmd := exec.Command("powershell", "-NoProfile", "-Command", psScript)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return nil
	}

	var metrics map[string]float64
	if err := json.Unmarshal(output, &metrics); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	return metrics
}
