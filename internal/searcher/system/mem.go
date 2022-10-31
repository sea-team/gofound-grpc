package system

import (
	gofoundpb "gofound-grpc/api/gen/v1"
	"runtime"

	"github.com/shirou/gopsutil/v3/mem"
)

// GetMemStat 内存信息
func GetMemInfo() (*gofoundpb.Memory, error) {
	info, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	//自身占用
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	self := GetUint64GB(memStat.Alloc)

	return &gofoundpb.Memory{
		Total:       GetUint64GB(info.Total),
		Used:        GetUint64GB(info.Used),
		Free:        GetUint64GB(info.Free),
		UsedPercent: GetPercent(info.UsedPercent),
		Self:        self,
	}, nil
}
