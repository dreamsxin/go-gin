package monitor

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

type MonitorInfo struct {
	Total         uint64
	Free          uint64
	Used          uint64
	UsedPercent   float64
	CpuPercent    []float64
	PhysicalCores int
	LogicalCores  int
	Process       []MonitorProcessInfo
}

type MonitorProcessInfo struct {
	Pid        int32
	Name       string
	Cmdline    string
	Cwd        string
	Status     []string
	NumThreads int32
	CpuPercent float64
	MemPercent float32
}

func Monitor() *MonitorInfo {
	monitorInfo := &MonitorInfo{}
	v, _ := mem.VirtualMemory()

	monitorInfo.Total = v.Total
	monitorInfo.Free = v.Free
	monitorInfo.Used = v.Used
	monitorInfo.UsedPercent = v.UsedPercent

	monitorInfo.CpuPercent, _ = cpu.Percent(time.Duration(time.Second), false)
	monitorInfo.PhysicalCores, _ = cpu.Counts(false)
	monitorInfo.LogicalCores, _ = cpu.Counts(true)

	processes, _ := process.Processes()
	for _, p := range processes {
		var processinfo MonitorProcessInfo
		processinfo.Pid = p.Pid
		processinfo.Name, _ = p.Name()
		processinfo.Cmdline, _ = p.Cmdline()
		processinfo.Cwd, _ = p.Cwd()
		processinfo.Status, _ = p.Status()

		processinfo.NumThreads, _ = p.NumThreads()
		processinfo.CpuPercent, _ = p.Percent(0)
		processinfo.MemPercent, _ = p.MemoryPercent()

		monitorInfo.Process = append(monitorInfo.Process, processinfo)
	}
	return monitorInfo
}
