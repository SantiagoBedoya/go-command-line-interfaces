package pkg

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/cpu"
)

func GetCpuInfo() []table.Row {
	cpuInfos, _ := cpu.Info()
	cpuInfo := cpuInfos[0]
	rows := []table.Row{
		{"CPU", "model", cpuInfo.Model},
		{"", "model_name", cpuInfo.ModelName},
		{"", "familiy", cpuInfo.Family},
		{"", "vendor_id", cpuInfo.VendorID},
		{"", "cache_size", cpuInfo.CacheSize},
		{"", "physical_id", cpuInfo.PhysicalID},
		{"", "cores", cpuInfo.Cores},
		{"", "frequency", cpuInfo.Mhz},
		{"", "cpu", cpuInfo.CPU},
	}
	return rows
}
