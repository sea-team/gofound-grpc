package system

import (
	"testing"
)

func TestMem(t *testing.T) {
	memory, err := GetMemInfo()
	if err != nil {
		t.Fatalf("cannot get memory status: %v", err)
	}

	t.Log(memory)
}
