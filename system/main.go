package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SystemInfo struct {
	CPUInfo    string `json:"cpu_info"`
	MemoryInfo string `json:"memory_info"`
}

func main() {
	systemInfo := SystemInfo{
		CPUInfo:    getCPUInfo(),
		MemoryInfo: getMemoryInfo(),
	}

	jsonData, err := json.MarshalIndent(systemInfo, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling system info:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func getCPUInfo() string {
	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error reading CPU info:", err)
		return ""
	}
	return string(data)
}

func getMemoryInfo() string {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println("Error reading memory info:", err)
		return ""
	}
	return string(data)
}
