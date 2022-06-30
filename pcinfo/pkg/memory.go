package pkg

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func GetMemoryInfo() []table.Row {
	parts, _ := disk.Partitions(true)
	virtualMemory, _ := mem.VirtualMemory()
	rows := []table.Row{
		{"Memory", "virtual_memory_total", humanize.Bytes(virtualMemory.Total)},
		{"", "virtual_memory_used", humanize.Bytes(virtualMemory.Used)},
		{"", "virtual_memory_free", humanize.Bytes(virtualMemory.Free)},
		// {"Memory", "total_memory", humanize.Bytes(memory.TotalMemory())},
		// {"", "free_memory", humanize.Bytes(memory.FreeMemory())},
	}
	for _, part := range parts {
		device := part.Mountpoint
		s, _ := disk.Usage(device)
		property := fmt.Sprintf("%s (%s)", part.Device, part.Fstype)
		if s != nil {
			value := fmt.Sprintf("Total=%s, Used=%s, Free=%s, Percentage=%2.f%%",
				humanize.Bytes(s.Total),
				humanize.Bytes(s.Used),
				humanize.Bytes(s.Free),
				s.UsedPercent,
			)
			rows = append(rows, table.Row{"", property, value})
		}
		// fmt.Println(part.String())
	}
	return rows
}
