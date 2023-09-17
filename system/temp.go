package system

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const systemTemperatureFile = "/sys/class/thermal/thermal_zone0/temp"

type SystemTemperature struct{}

// GetDisplayValueForSystemTemperature returns a string value for the specified
func (a SystemTemperature) GetDisplayValueForSystemTemperature(thermalZone string) string {
	temperature, err := a.InCByZone(thermalZone)
	if err != nil || temperature == nil {
		return "XX"
	}
	return fmt.Sprintf("%v", *temperature/1000)
}

// InCByZone returns the systems temperature in Celsius for the thermal zone specified.
func (SystemTemperature) InCByZone(thermalZone string) (*int64, error) {
	if runtime.GOOS == "linux" && runtime.GOARCH == "arm64" {
		log.Printf("found runtime as `%s` and `%s`", runtime.GOOS, runtime.GOARCH)
		tempData, err := os.ReadFile(systemTemperatureFile)
		if err != nil {
			return nil, err
		}

		var cleanedStringTempData = strings.TrimSpace(string(tempData))
		log.Printf("found tempData as `%s`", cleanedStringTempData)

		intValue, err := strconv.ParseInt(cleanedStringTempData, 10, 64)
		if err != nil {
			return nil, err
		}
		return &intValue, nil
	}

	return nil, fmt.Errorf("unsupported runtime")
}
