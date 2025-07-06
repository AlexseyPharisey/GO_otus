package settings

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func DiskMemoryMetrics(diskName string) map[string]float64 {
	psScript := fmt.Sprintf(`
		$disk="%s:";
		$d=Get-CimInstance Win32_LogicalDisk | Where-Object {
			$_.DeviceID -eq $disk
		};
		$sizeMB=$d.Size/1MB;
		$freeMB=$d.FreeSpace/1MB;
		$usedPct=[math]::Round((($sizeMB - $freeMB)/$sizeMB)*100,2);
		$inodes=(Get-ChildItem "$disk\" -Force -ErrorAction SilentlyContinue | Measure-Object).Count;
		@{
			UsedPercent=$usedPct;
			UsedInodePercent=100
		} | ConvertTo-Json
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
