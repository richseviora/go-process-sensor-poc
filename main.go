package main

import (
	"context"
	"os"
	"os/signal"
	"process-sensor-poc/src"
	"sync"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go src.RecurringTimer(ctx, wg)
	<-c
	cancel()
	wg.Wait()
}
