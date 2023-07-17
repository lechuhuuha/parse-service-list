package parse_service_list

import (
	"fmt"
	"regexp"
	"strings"
)

type SystemdItems struct {
	Name        string `json:"serviceName"`
	Loaded      string `json:"loaded"`
	State       string `json:"state"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func ParseSystemdOutput(bytestream []byte) ([]SystemdItems, error) {
	var serviceItemsList []SystemdItems
	bsString := string(bytestream)

	lines := strings.Split(bsString, "\n")

	if len(lines) <= 1 {
		return nil, fmt.Errorf("unable to parse bytestream []byte: insufficient lines")
	}

	for k, v := range lines {
		if k == 0 { //header ?? skip
			continue
		}

		re := regexp.MustCompile(`\s(.+\.service)\s+([A-z]+)\s+([A-z]+)\s+([A-z]+)\s+(.+)`)
		segments := re.FindAllStringSubmatch(v, -1)

		if len(segments) > 0 {
			si := SystemdItems{
				Name:        strings.Trim(segments[0][1], " "),
				Loaded:      strings.Trim(segments[0][2], " "),
				State:       strings.Trim(segments[0][3], " "),
				Status:      strings.Trim(segments[0][4], " "),
				Description: strings.Trim(segments[0][5], " "),
			}
			serviceItemsList = append(serviceItemsList, si)
		}
	}

	return serviceItemsList, nil
}

type ProcessStatusItems struct {
	PID     int     `json:"pid"`
	User    string  `json:"user"`
	Nice    int     `json:"ni"`
	CPU     float64 `json:"cpu"`
	Memory  float64 `json:"mem"`
	Command string  `json:"command"`
}

func ParsePSOutput(bytestream []byte) ([]ProcessStatusItems, error) {
	var processes []ProcessStatusItems
	bsString := string(bytestream)

	lines := strings.Split(bsString, "\n")

	if len(lines) <= 1 {
		return nil, fmt.Errorf("unable to parse bytestream []byte: insufficient lines")
	}

	for k, v := range lines {
		if k == 0 { // header - skip
			continue
		}

		fields := strings.Fields(v)

		if len(fields) < 6 {
			continue // Skip invalid lines
		}

		pid := 0
		fmt.Sscanf(fields[0], "%d", &pid)

		user := fields[1]

		ni := 0
		fmt.Sscanf(fields[2], "%d", &ni)

		cpu := 0.0
		fmt.Sscanf(fields[3], "%f", &cpu)

		mem := 0.0
		fmt.Sscanf(fields[4], "%f", &mem)

		command := fields[5]

		process := ProcessStatusItems{
			PID:     pid,
			User:    user,
			Nice:    ni,
			CPU:     cpu,
			Memory:  mem,
			Command: command,
		}

		processes = append(processes, process)
	}

	return processes, nil
}

type DiskUsages struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Perc   string `json:"pcent"`
	Used   string `json:"used"`
	Total  string `json:"itotal"`
}

func ParseDiskUsage(output []byte) ([]DiskUsages, error) {
	lines := strings.Split(string(output), "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("no disk usage data found")

	}
	diskUsages := make([]DiskUsages, 0)

	for i := 1; i < len(lines); i++ {
		fields := strings.Fields(lines[i])
		if len(fields) >= 5 {
			diskUsage := DiskUsages{
				Source: fields[0],
				Target: fields[1],
				Perc:   fields[2],
				Used:   fields[3],
				Total:  fields[4],
			}
			diskUsages = append(diskUsages, diskUsage)
		}
	}
	if len(diskUsages) == 0 {
		return nil, fmt.Errorf("no disk usage data found")
	}

	return diskUsages, nil
}
