package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	go scheduleTest()
	time.Sleep(60 * time.Second)
}

type myjob struct {

}
func (job myjob) Run() {
	fmt.Println("scheduler myjob")
}
func scheduleTest() (error) {
	c := cron.New()
	myjob := new(myjob)
	//err := c.AddFunc("*/5 * * * * *", func() {
	//	fmt.Println("schedule func")
	//})
	err := c.AddJob("*/5 * * * * *", myjob)
	if err != nil {
		return err
	}
	c.Start()
	return nil
}