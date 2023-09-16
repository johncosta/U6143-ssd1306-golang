package system

import (
	"os"
	"runtime"
	"strconv"
)

const systemTemperatureFile = "/sys/class/thermal/thermal_zone0/temp"

type SystemTemperature struct{}

// GetDisplayValueForSystemTemperature returns a string value for the specified
func (a SystemTemperature) GetDisplayValueForSystemTemperature(thermalZone string) string {
	temperature, err := a.InCByZone(thermalZone)
	if err != nil || temperature == nil {
		return "XX"
	}
	return string(*temperature)
}

// InCByZone returns the systems temperature in Celsius for the thermal zone specified.
func (SystemTemperature) InCByZone(thermalZone string) (tempInC *int, err error) {
	if runtime.GOOS == "linux" && runtime.GOARCH == "arm64" {
		tempData, err := os.ReadFile(systemTemperatureFile)
		if err != nil {
			return nil, err
		}

		*tempInC, err = strconv.Atoi(string(tempData))
		if err != nil {
			return nil, err
		}
	}

	return tempInC, nil
}
