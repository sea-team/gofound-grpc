package system

import (
	"testing"
)

func TestDisk(t *testing.T) {
	disk, err := GetDiskInfo()
	if err != nil {
		t.Fatalf("cannot get disk status: %v", err)
	}

	t.Log(disk)
}
