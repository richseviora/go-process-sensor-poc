package src

import (
	"context"
	"fmt"
	"sync"
	"time"
)

import "github.com/mitchellh/go-ps"

func getProcessInfo() ([]ps.Process, error) {
	return ps.Processes()
}

func QueryProcesses() {
	printProcessInfo()
}

func RecurringTimer(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
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
