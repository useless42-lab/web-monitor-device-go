package utils

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

//获取cpu方面的信息
func GetCpuInfo() (cpuUser float64, cpuSystem float64, cpuIdle float64, cpuPercent float64) {
	// 	cpu // 具体cpu名称
	// user // 用户已使用cpu数值
	// system // 系统已使用cpu数值
	// idle // 闲置的cpu数值
	// .Total // 获取总cpu数值
	res, err := cpu.Times(false) // false是展示全部总和 true是分布展示
	if err != nil {
		fmt.Println(err)
	}
	percent, _ := cpu.Percent(time.Second, false)
	return res[0].User, res[0].System, res[0].Idle, percent[0]
}

//获取内存方面的信息
func GetMemoryInfo() (memoryTotal uint64, memoryAvailable uint64, memoryUsed uint64, memoryUsedPercent float64) {
	// 	total // 内存大小
	// available // 闲置可用内存
	// used // 已使用内存
	// usedPercent // 已使用百分比
	v, _ := mem.VirtualMemory()
	return v.Total, v.Available, v.Used, v.UsedPercent
}

// disk info
func GetDiskInfo() (diskTotal uint64, diskFree uint64, diskUsed uint64, diskUsedPercent float64) {
	parts, err := disk.Partitions(false)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.Total, diskInfo.Free, diskInfo.Used, diskInfo.UsedPercent
}

func GetNetInfo() (netSent uint64, netRecv uint64) {
	// net 发送 // BytesSent
	// net 接收 // BytesRecv
	info, _ := net.IOCounters(false)
	return info[0].BytesSent, info[0].BytesRecv
}
