package pkg

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/host"
)

func GetHostInfo() []table.Row {
	hostInfo, _ := host.Info()
	rows := []table.Row{
		{"Host", "hostname", hostInfo.Hostname},
		{"", "host_id", hostInfo.HostID},
		{"", "uptime", hostInfo.Uptime},
		{"", "number_of_processes_running", hostInfo.Procs},
		{"", "OS", hostInfo.OS},
		{"", "platform", hostInfo.Platform},
	}

	return rows
}
