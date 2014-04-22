// +build linux freebsd

package gopsutil

import (
	"os"
	"syscall"
)

func HostInfo() (HostInfoStat, error) {
	ret := HostInfoStat{}

	hostname, err := os.Hostname()
	ret.Hostname = hostname
	if err != nil {
		return ret, err
	}

	sysinfo := &syscall.Sysinfo_t{}
	if err := syscall.Sysinfo(sysinfo); err != nil {
		return ret, err
	}
	ret.Uptime = sysinfo.Uptime

	return ret, nil
}