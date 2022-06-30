package pkg

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/net"
)

func GetMacInfo() []table.Row {
	interfStat, _ := net.Interfaces()
	rows := []table.Row{}

	for i, inter := range interfStat {
		if i == 0 {
			rows = append(rows,
				table.Row{"MAC", fmt.Sprintf("%d. %s", i+1, "interface_name"), inter.Name},
				table.Row{"", "mac_address", inter.HardwareAddr},
				table.Row{"", "interface_behavior_or_flags", strings.Join(inter.Flags, ",")},
			)
		} else {
			rows = append(rows,
				table.Row{"", fmt.Sprintf("%d. %s", i+1, "interface_name"), inter.Name},
				table.Row{"", "mac_address", inter.HardwareAddr},
				table.Row{"", "interface_behavior_or_flags", strings.Join(inter.Flags, ",")},
				table.Row{"", "IPv6_or_IPv4_addresses", getAddresses(inter.Addrs)},
			)
		}
	}

	return rows
}

func getAddresses(addrs []net.InterfaceAddr) string {
	str := []string{}
	for _, add := range addrs {
		str = append(str, add.String())
	}
	return strings.Join(str, ",")
}
