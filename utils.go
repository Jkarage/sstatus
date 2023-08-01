package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Variables to store previous values for calculating CPU usage
var prevIdle, prevTotal int64

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

// Function to calculate the CPU usage percentage
func getCPUPercentage() float64 {
	statFile, err := os.Open("/proc/stat")
	if err != nil {
		fmt.Println("Error opening /proc/stat:", err)
		return 0.0
	}
	defer statFile.Close()

	scanner := bufio.NewScanner(statFile)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) > 0 && fields[0] == "cpu" {
			var total, idle int64
			for i := 1; i < len(fields); i++ {
				value, err := parseField(fields[i])
				if err != nil {
					fmt.Println("Error parsing CPU usage:", err)
					return 0.0
				}

				total += value
				if i == 4 { // The 5th field is "idle"
					idle = value
				}
			}

			// Calculate CPU usage percentage
			idleDiff := float64(idle - prevIdle)
			totalDiff := float64(total - prevTotal)
			cpuUsage := 100 * (totalDiff - idleDiff) / totalDiff

			// Update previous values for the next iteration
			prevIdle = idle
			prevTotal = total

			return cpuUsage
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading /proc/stat:", err)
	}

	return 0.0
}

// Helper function to parse integer value from string
func parseField(field string) (int64, error) {
	return strconv.ParseInt(field, 10, 64)
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
