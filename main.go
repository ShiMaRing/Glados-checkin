package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
	"os"
	"time"
)

func main() {

	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE, 7777)
	if err != nil {
		panic("open log file fail")
	}
	execTime := viper.GetString("execTime")
	if execTime == "" {
		panic("Illegal execution time")
	}
	s := gocron.NewScheduler(time.UTC)

	job, err := s.Every(1).Days().At(execTime).Do(checkin)

	if err != nil {
		_, _ = fmt.Fprintln(file, fmt.Sprintf("Job: %v, Error: %v", job, err))
	}

	s.StartBlocking()

}
