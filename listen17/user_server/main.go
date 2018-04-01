package main

import (
	"time"

	"github.com/pingguoxueyuan/gostudy/logger"
)

func initLogger(name, logPath, logName string, level string) (err error) {
	m := make(map[string]string, 8)
	m["log_path"] = logPath
	m["log_name"] = "user_server"
	m["log_level"] = level
	err = logger.InitLogger(name, m)
	if err != nil {
		return
	}

	logger.Debug("init logger success")
	return
}

func Run() {
	for {
		logger.Debug("user server is running")
		time.Sleep(time.Second)
	}
}

func main() {
	initLogger("file", "c:/logs/", "user_server", "debug")
	Run()
	return
}
