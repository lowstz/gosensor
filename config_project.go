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
			Port int `json:"port",omitempty`
		} `json:"detail"`
		Frequency string `json:"frequency"`
		Type      string `json:"type"`
	} `json:"monitor"`
	Name   string `json:"name"`
	Notify struct {
		Email struct {
			Content struct {
			} `json:"content"`
			MailPassword string        `json:"mail_password"`
			MailUser     string        `json:"mail_user"`
			SendToList   []interface{} `json:"send_to_list"`
			SmtpPort     int64         `json:"smtp_port"`
			SmtpServer   string        `json:"smtp_server"`
			Subject      string        `json:"subject"`
		} `json:"email"`
		Hipchat struct {
			Content struct {
				Offline string `json:"offline"`
				Online  string `json:"online"`
			} `json:"content"`
			From  string `json:"from"`
			To    string `json:"to"`
			Token string `json:"token"`
		} `json:"hipchat"`
	} `json:"notify"`
}

type ProjectConfigParser struct{}

var (
	configPath string
)

func init() {
	flag.StringVar((*string)(&configPath), "rc", "./monitor.json", "project config file path")
}

func (p ProjectConfigParser) Parse(path string) ProjectConfig {
	exists, err := files.IsFileExist(files.AbsPath(path))
	if err != nil {
		errl.Println(showCallerName(), err)
	}
	if !exists {
		errl.Println(showCallerName(), "cannot access", path, ": No such file or directory")
		os.Exit(1)
	}

	file, _ := os.Open(files.AbsPath(path))
	defer file.Close()
	decoder := json.NewDecoder(file)
	projectConfig := ProjectConfig{}
	err = decoder.Decode(&projectConfig)
	if err != nil {
		errl.Println("parse err")
	}
	return projectConfig
}
