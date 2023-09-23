package system

import (
	"fmt"
	"github.com/pbnjay/memory"
)

type Memory struct{}

func (Memory) GetDisplayValueForSystemMemory() string {
	return fmt.Sprintf("Free %vmb/ Total %vmb", memory.FreeMemory()/100000, memory.TotalMemory()/100000)
}
