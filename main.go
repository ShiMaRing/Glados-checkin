package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
	"os"
	"time"
)

var file *os.File

func init() {
	var err error
	file, err = os.OpenFile("log.log", os.O_APPEND|os.O_CREATE, 7777)
	if err != nil {
		panic(err)
	}
}
func main() {
	execTime := viper.GetString("execTime")
	if execTime == "" {
		panic("Illegal execution time")
	}
	s := gocron.NewScheduler(time.UTC)
	job, err := s.Every(time.Hour * 16).Do(checkin)
	if err != nil {
		sprintf := fmt.Sprintf("Job: %v, Error: %v", job, err)
		Log(sprintf)
	}
	s.StartBlocking()
	_, _ = file.WriteString("Finished")
	if err != nil {
		return
	}
	file.Close()
}

func Log(message string) {
	_, _ = fmt.Fprintln(file, message)
	_ = file.Sync()
}
