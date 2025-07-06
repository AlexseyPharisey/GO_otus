package settings

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func GetDiskName() []string {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"Get-CimInstance Win32_LogicalDisk -Filter \"DriveType=3\""+
			" | Select-Object -ExpandProperty DeviceID | ForEach-Object { $_.TrimEnd(\":\") } | ConvertTo-Json\n",
	)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return nil
	}

	var metrics []string
	if err := json.Unmarshal(output, &metrics); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	return metrics
}
