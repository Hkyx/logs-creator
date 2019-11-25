package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"log"
	"os"
	"time"
)

func task() {
	fmt.Println(string(`{"adrien": "12345", "adrien2": "abc", "time": "28/Jul/2006:10:22:04 -0300"}`))
}

func taskfile() {
	pathLog := "/var/log/genlog.log"

	f, err := os.OpenFile(pathLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	t := time.Now()
	p := t.Format("2006-01-02T15:04:05.999999-0300")

	test, err := f.WriteString(string(`[{adrien": "12345", "adrien2": "abc", "time": "` + string(p) + `"` + "}]\n"))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(test)
}

func taskfilejson() {
	pathLog := "/var/log/1.json"

	f, err := os.OpenFile(pathLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	t := time.Now()
	p := t.Format("2006-01-02T15:04:05.999999-0300")

	test, err := f.WriteString(string(`[{log1": "12345", "logfile": "def", "time": "` + string(p) + `"` + "}]\n"))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(test)
}

func main() {
	// Do jobs without params
	gocron.Every(1).Second().Do(taskfile)
	gocron.Every(1).Second().Do(taskfile)
	// remove, clear and next_run

	// gocron.Remove(task)
	// gocron.Clear()

	// function Start start all the pending jobs
	<-gocron.Start()
}
