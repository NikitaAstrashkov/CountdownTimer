package main

import (
	"fmt"
	"time"
)

func main() {
	v, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	fmt.Println(v.Unix(), time.Now().Unix())
	go func() {
		sleepDuration := time.Duration(v.Unix()-time.Now().Unix()) * time.Second
		fmt.Println(sleepDuration)
		time.Sleep(sleepDuration)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			timeTillNY := v.Unix() - t.Unix()
			fmt.Printf("Time untill NY: %d %d %d %d\n", timeTillNY/3600/24, timeTillNY/3600%24, timeTillNY/60%24%60, timeTillNY%60)
		}
	}
}
