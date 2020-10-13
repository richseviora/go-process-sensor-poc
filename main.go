package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

import "github.com/mitchellh/go-ps"

func getProcessInfo() ([]ps.Process, error) {
	return ps.Processes()
}

func recurringTimer(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	fmt.Println("Starting Timer")
	fmt.Println("Starting Goroutine")
	for {
		fmt.Println(time.Now().String())
		// Select only for channels
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-ticker.C:
			printProcessInfo()
		}
	}
}

func printProcessInfo() {
	processes, err := getProcessInfo()
	if err != nil {
		return
	}
	for _, p := range processes {
		fmt.Printf("Name %s PID %d \n", p.Executable(), p.Pid())
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go recurringTimer(ctx, wg)
	<-c
	cancel()
	wg.Wait()
}
