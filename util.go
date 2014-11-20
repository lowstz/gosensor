package main

import (
	"net"
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

func getLocalIPAddress() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		errl.Println(showCallerName(), err)
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				debug.Println(showCallerName(), "ip address is: ", ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return "ip not found"
}
