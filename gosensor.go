package main

import (
	"flag"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Gosensor struct{}

func (g Gosensor) CheckServiceAliveWithPort(port int) bool {
	cmd := "lsof -i:" + strconv.Itoa(port) + " | grep -v COMMAND | wc -l"
	debug.Println(showCallerName(), cmd)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		errl.Println(showCallerName(), err)
	}
	debug.Println(showCallerName(), "command output: ", out)
	result := strings.TrimSpace(string(out))
	debug.Println(showCallerName(), "result: ", result)
	status, err := strconv.Atoi(result)
	if err != nil {
		errl.Println(showCallerName(), err)
	}
	debug.Println(showCallerName(), status)
	if status > 0 {
		return true
	}
	return false
}

func main() {
	flag.Parse()
	initLog()

	var gosensor Gosensor
	var parser ProjectConfigParser

	config := parser.Parse(configPath)

	var wg sync.WaitGroup
	freq, err := strconv.Atoi(config.Monitor.Frequency)
	info.Println(showCallerName()+"freq: ", freq)
	if err != nil {
		errl.Println(showCallerName(), "strconv err", err)
	}
	ticker := time.NewTicker(time.Second * time.Duration(freq))

	for time := range ticker.C {
		debug.Println(showCallerName(), time)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if gosensor.CheckServiceAliveWithPort(config.Monitor.Detail.Port) {
				info.Println(showCallerName(), getHostName(), config.Name, "is alive.")
			} else {
				errl.Println(showCallerName(), getHostName(), config.Name, "is down!!!")
			}
		}()
		wg.Wait()
	}
}
