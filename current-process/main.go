package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type ProcessInfo struct {
	FileDescriptorsCount int    `json:"file_descriptors_count"`
	MemoryUsage          string `json:"memory_usage"`
	ExecutablePath       string `json:"executable_path"`
}

func main() {
	processInfo := ProcessInfo{
		FileDescriptorsCount: getFileDescriptorsCount(),
		MemoryUsage:          getMemoryUsage(),
		ExecutablePath:       getExecutablePath(),
	}

	jsonData, err := json.MarshalIndent(processInfo, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling process info:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func getFileDescriptorsCount() int {
	files, err := os.ReadDir("/proc/self/fd")
	if err != nil {
		fmt.Println("Error reading file descriptors:", err)
		return -1
	}
	return len(files)
}

func getMemoryUsage() string {
	data, err := os.ReadFile("/proc/self/status")
	if err != nil {
		fmt.Println("Error reading memory usage:", err)
		return ""
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "VmSize:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "VmSize:"))
		}
	}
	return ""
}

func getExecutablePath() string {
	path, err := os.Readlink("/proc/self/exe")
	if err != nil {
		fmt.Println("Error reading executable path:", err)
		return ""
	}
	return path
}
