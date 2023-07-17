package parse_service_list

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func collectServiceInfo() ([]SystemdItems, error) {
	cmd := exec.Command("systemctl", "list-units", "--type=service")
	outputData, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	serviceItemsList, err := ParseSystemdOutput(outputData)
	if err != nil {
		return serviceItemsList, err
	}

	return serviceItemsList, nil
}

func TestParseSystemdOutput(t *testing.T) {
	serviceItemsList, err := collectServiceInfo()
	if err != nil {
		t.Errorf("Error collecting service info: %v", err)
	}

	// Perform assertions on the serviceItemsList
	if len(serviceItemsList) == 0 {
		t.Error("Expected non-empty serviceItemsList, but got empty")
	}

	// Define the expected struct fields
	expectedFields := []string{"Name", "Loaded", "State", "Status", "Description"}

	// Check each element in serviceItemsList for the correct struct fields
	for _, service := range serviceItemsList {
		serviceType := reflect.TypeOf(service)
		for _, field := range expectedFields {
			_, found := serviceType.FieldByName(field)
			if !found {
				t.Errorf("Expected field %q not found in serviceItemsList element: %v", field, service)
			}
		}
		fmt.Printf("Service Name: %s\nLoaded: %s\nState: %s\nStatus: %s\nDescription: %s\n\n", service.Name, service.Loaded, service.State, service.Status, service.Description)
	}

}

func collectPsInfo() ([]ProcessStatusItems, error) {
	cmd := exec.Command("ps", "-eo", "pid,user,ni,%cpu,%mem,args", "--sort=-%cpu,-%mem")
	outputData, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	serviceItemsList, err := ParsePSOutput(outputData)
	if err != nil {
		return serviceItemsList, err
	}

	return serviceItemsList, nil
}

func TestParsePSOutput(t *testing.T) {
	serviceItemsList, err := collectPsInfo()
	if err != nil {
		t.Errorf("Error collecting service info: %v", err)
	}

	// Perform assertions on the serviceItemsList
	if len(serviceItemsList) == 0 {
		t.Error("Expected non-empty serviceItemsList, but got empty")
	}

	for _, process := range serviceItemsList {
		serviceType := reflect.TypeOf(process)
		expectedFields := []string{"PID", "User", "Nice", "CPU", "Memory", "Command"}
		for _, field := range expectedFields {
			_, found := serviceType.FieldByName(field)
			if !found {
				t.Errorf("Expected field %q not found in SystemdItems element: %v\n", field, process)
			}
		}

		fmt.Printf("PID: %d User: %s Nice: %d CPU: %.1f Memory: %.1f Command: %s\n",
			process.PID, process.User, process.Nice, process.CPU, process.Memory, process.Command)

	}
}

func TestCollectPsInfoAndJsonMarshal(t *testing.T) {

	serviceItemsList, err := collectPsInfo()
	if err != nil {
		t.Errorf("Error collecting service info: %v", err)
	}

	// Perform assertions on the serviceItemsList
	if len(serviceItemsList) == 0 {
		t.Error("Expected non-empty serviceItemsList, but got empty")
	}
	// Convert the slice to JSON
	jsonData, err := json.Marshal(serviceItemsList)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON
	fmt.Println(string(jsonData))
}

func TestCollectServiceInfoAndJsonMarshal(t *testing.T) {

	serviceItemsList, err := collectServiceInfo()
	if err != nil {
		t.Errorf("Error collecting service info: %v", err)
	}

	// Perform assertions on the serviceItemsList
	if len(serviceItemsList) == 0 {
		t.Error("Expected non-empty serviceItemsList, but got empty")
	}
	// Convert the slice to JSON
	jsonData, err := json.Marshal(serviceItemsList)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON
	fmt.Println(string(jsonData))
}
