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
	oldProcesses, err := getProcessInfo()
	if err != nil {
		return
	}
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-ticker.C:
			newProcesses, err := getProcessInfo()
			if err != nil {
				return
			}
			printChanges(GetChanges(oldProcesses, newProcesses))
			oldProcesses = newProcesses
		}
	}
}

func printChanges(dr DiffResult) {
	fmt.Println("New Processes")
	for _, p := range dr.Added() {
		printProcessLine(p)
	}
	fmt.Println("Dropped Processes")
	for _, p := range dr.Removed() {
		printProcessLine(p)
	}
}

func printProcessLine(p ps.Process) {
	fmt.Printf("Name %s PID %d \n", p.Executable(), p.Pid())
}

func printProcessInfo() {
	processes, err := getProcessInfo()
	if err != nil {
		return
	}
	for _, p := range processes {
		printProcessLine(p)
	}
}
