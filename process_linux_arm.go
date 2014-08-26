// +build linux
// +build arm

package gopsutil

const (
	ClockTicks = 100  // C.sysconf(C._SC_CLK_TCK)
	PageSize   = 4096 // C.sysconf(C._SC_PAGE_SIZE)
)