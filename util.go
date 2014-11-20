package main

import (
	"os"
	"runtime"
)

func showCallerName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		errl.Println("rumtime.Caller() failed\n")
		return "unkown caller"
	}
	return runtime.FuncForPC(pc).Name()
}

func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		errl.Println(showCallerName(), err)
		return "Unknown Hostname"
	} else {
		debug.Println(hostname)
		return hostname
	}
}
