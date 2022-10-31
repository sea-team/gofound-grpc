package system

import (
	"testing"
)

func TestCPU(t *testing.T) {
	cpu, err := GetCPUInfo()
	if err != nil {
		t.Fatalf("cannot get cup status: %v", err)
	}

	t.Log(cpu)
}
