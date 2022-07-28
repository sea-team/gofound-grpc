package system

import (
	gofoundpb "gofound-grpc/api/gen/v1"

	"github.com/shirou/gopsutil/v3/disk"
)

// GetDiskInfo 磁盘信息
func GetDiskInfo() (*gofoundpb.Disk, error) {
	parts, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	if err != nil {
		return nil, err
	}

	return &gofoundpb.Disk{
		Path:        diskInfo.Path,
		Total:       GetUint64GB(diskInfo.Total),
		Free:        GetUint64GB(diskInfo.Free),
		Used:        GetUint64GB(diskInfo.Used),
		UsedPercent: GetPercent(diskInfo.UsedPercent),
		FsType:      diskInfo.Fstype,
	}, nil
}
