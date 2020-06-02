package main

import (
	"basego/runner"
	"log"
	"os"
	"time"
)

const timeout = 4 * time.Second

func main() {

	log.Println("start working...")

	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		}
	}

	log.Println("Process ended...")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Proceser - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
