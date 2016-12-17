package main

import (
	"syscall"

	"github.com/cihub/seelog"
)

func main() {
	seelog.Info("Hello,I am stock sniper!Biu Biu Biu ~")

	var rlimit syscall.Rlimit
	syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)

	seelog.Info("rlimit is ", rlimit)
	rlimit.Cur = 65536
	rlimit.Max = 65536
	syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit)

	//UserAPPCode
	// realtimelib.RealtimeOneStock("300345")

	seelog.Flush()
}
