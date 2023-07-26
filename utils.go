package main

import (
	"fmt"
	"os"
	"strings"
)

type MemoryInfo struct {
	Total     uint64
	Free      uint64
	Available uint64
}

func getMemoryStatus() (*MemoryInfo, error) {
	info, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	return parseMemInfo(string(info)), nil
}

func parseMemInfo(meminfo string) *MemoryInfo {
	m := MemoryInfo{}
	lines := strings.Split(meminfo, "\n")

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			switch fields[0] {
			case "MemTotal:":
				fmt.Sscanf(fields[1], "%d", &m.Total)
			case "MemFree:":
				fmt.Sscanf(fields[1], "%d", &m.Free)
			case "MemAvailable:":
				fmt.Sscanf(fields[1], "%d", &m.Available)
			}
		}
	}

	return &m
}

func getCPUStats() (float64, float64, error) {
	stats, err := os.ReadFile("/proc/stat")
	if err != nil {
		return 0, 0, err
	}

	lines := strings.Split(string(stats), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 0 && fields[0] == "cpu" {
			var totalTime, idleTime float64
			for i := 1; i < len(fields); i++ {
				var val float64
				fmt.Sscanf(fields[i], "%f", &val)
				totalTime += val
				if i == 4 {
					idleTime = val
				}
			}
			return totalTime - idleTime, idleTime, nil
		}

	}

	return 0, 0, fmt.Errorf("couldn't find cpu stats in /proc/stats")
}
