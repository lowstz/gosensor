package main

import (
	"flag"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	version = "0.0.1"
)

type Gosensor struct{}

func (g Gosensor) CheckServiceAliveWithPort(port int) bool {
	cmd := "lsof -i:" + strconv.Itoa(port) + " | grep LISTEN | wc -l"
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
	downStatusCounter := 0
	for time := range ticker.C {
		debug.Println(showCallerName(), time)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if gosensor.CheckServiceAliveWithPort(config.Monitor.Detail.Port) {
				info.Println(showCallerName(), getHostName(), config.Name, "is alive.")
				if downStatusCounter != 0 {
					notifReq := Hipchat{
						Token: config.Notify.Hipchat.Token,
						From:  config.Notify.Hipchat.From,
						To:    config.Notify.Hipchat.To,
						Color: "green",
						Content: "INFO:  {" + getHostName() + ":" + getLocalIPAddress() +
							"} [" + config.Name + "] is back to alive \n" +
							config.Notify.Hipchat.Content.Online,
					}
					SendNotification(notifReq)
					downStatusCounter = 0
				}
			} else {
				errl.Println(showCallerName(), getHostName(), config.Name, "is down!!!")
				notifReq := Hipchat{
					Token: config.Notify.Hipchat.Token,
					From:  config.Notify.Hipchat.From,
					To:    config.Notify.Hipchat.To,
					Color: "red",
					Content: "ERROR: {" + getHostName() + ":" + getLocalIPAddress() +
						"} [" + config.Name + "] is down!!! \n" +
						config.Notify.Hipchat.Content.Offline,
				}
				SendNotification(notifReq)
				downStatusCounter++
			}
		}()
		wg.Wait()
	}
}
