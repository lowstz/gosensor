package main

import (
	"encoding/json"
	"flag"
	"github.com/magicshui/goutils/files"
	"os"
)

type ProjectConfig struct {
	Description string `json:"description"`
	Github      string `json:"github"`
	Master      struct {
		Ip string `json:"ip"`
	} `json:"master"`
	Monitor struct {
		Detail struct {
			Port    int    `json:"port",omitempty`
			PidFile string `json:"pidfile,omitempty"`
		} `json:"detail"`
		Frequency string `json:"frequency"`
		Type      string `json:"type"`
	} `json:"monitor"`
	Name   string `json:"name"`
	Notify struct {
		Detail struct {
			Token   string `json:token`
			Content string `json:"content"`
			From    string `json:"from"`
			To      string `json:"to"`
		} `json:"detail"`
		Type string `json:"type"`
	} `json:"notify"`
}

type ProjectConfigParser struct{}

var (
	configPath string
)

func init() {
	flag.StringVar((*string)(&configPath), "configPath", "./monitor.json", "project config file path")
}

func (p ProjectConfigParser) Parse(path string) ProjectConfig {

	file, _ := os.Open(files.AbsPath(path))
	defer file.Close()
	decoder := json.NewDecoder(file)
	projectConfig := ProjectConfig{}
	err := decoder.Decode(&projectConfig)
	if err != nil {
		errl.Println("parse err")
	}
	return projectConfig
}
