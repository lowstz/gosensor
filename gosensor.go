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

var gosensor Gosensor

func (g Gosensor) checkServiceAliveWithPort(port int) bool {
	cmd := "lsof -i:" + strconv.Itoa(port) + " | grep -v COMMAND | wc -l"
	debug.Println(getFuncName(), cmd)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		errl.Println(getFuncName(), err)
	}
	debug.Println(getFuncName(), "command output: ", out)
	result := strings.TrimSpace(string(out))
	debug.Println(getFuncName(), "result: ", result)
	status, err := strconv.Atoi(result)
	if err != nil {
		errl.Println(getFuncName(), err)
	}
	debug.Println(getFuncName(), status)
	if status > 0 {
		return true
	}
	return false
}

func main() {
	flag.Parse()
	initLog()

	var parser ProjectConfigParser

	config := parser.Parse(configPath)

	var wg sync.WaitGroup
	freq, err := strconv.Atoi(config.Monitor.Frequency)
	info.Println(getFuncName()+"freq: ", freq)
	if err != nil {
		errl.Println(getFuncName(), "strconv err", err)
	}
	ticker := time.NewTicker(time.Second * time.Duration(freq))

	for t := range ticker.C {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if gosensor.checkServiceAliveWithPort(config.Monitor.Detail.Port) {
				info.Println(getFuncName(), "Service work OK!", t)
			} else {
				errl.Println(getFuncName(), "Service is down!!!", t)
			}
		}()
		wg.Wait()
	}
}
