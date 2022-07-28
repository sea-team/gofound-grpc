package system

import (
	gofoundpb "gofound-grpc/api/gen/v1"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

// GetCPUInfo CPU信息
func GetCPUInfo() (*gofoundpb.CPU, error) {

	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	info, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	return &gofoundpb.CPU{
		UsedPercent: GetPercent(percent[0]),
		Cores:       int32(runtime.NumCPU()),
		ModelName:   info[0].ModelName,
	}, nil
}
