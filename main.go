package main

import (
	"WebMonitorDevice/jobs"
	"fmt"

	"github.com/bamzi/jobrunner"
)

func main() {
	jobrunner.Start()
	jobrunner.Schedule("@every 60s", jobs.CheckServerJob{})

	var str string
	fmt.Scan(&str)
}
