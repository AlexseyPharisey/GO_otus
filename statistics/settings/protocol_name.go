package settings

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func GetProtocolName() []string {
	cmd := exec.Command(
		"powershell",
		"-Command",
		`$protocols=@();
		if(Get-NetTCPConnection -EA SilentlyContinue){
			$protocols+='TCP'
		};
		if(Get-NetUDPEndpoint -EA SilentlyContinue){
			$protocols+='UDP'
		};
		if(Get-NetFirewallRule -DisplayName '*ICMP*' -EA SilentlyContinue){
			$protocols+='ICMP'
		};
		$protocols|Sort-Object|Get-Unique|ConvertTo-Json`,
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
