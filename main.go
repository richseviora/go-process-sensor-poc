package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type ProcessInfo struct {
	name      string
	arguments string
	pid       int32
}

func getProcessInfo() *bytes.Buffer {
	cmd := exec.Command("ps")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return &bytes.Buffer{}
	}
	return &out
}

func parseProcessLine(line string) ProcessInfo {
	stringPid := strings.TrimSpace(line[0:6])
	pid, err := strconv.ParseInt(stringPid, 10, 32)
	if err != nil {
		fmt.Println(err)
		return ProcessInfo{}
	}
	return struct {
		name      string
		arguments string
		pid       int32
	}{
		name: strings.TrimSpace(line[27:]),
		// TODO: Figure out how to parse arguments from command name.
		arguments: "",
		pid:       int32(pid),
	}
}

func parseProcessInfo(b *bytes.Buffer) []ProcessInfo {
	var results []ProcessInfo
	var err error
	var line string
	// Drop first line
	line, err = b.ReadString('\n')
	fmt.Printf("First Line %s \n", line)
	for {
		line, err = b.ReadString('\n')
		fmt.Printf("Parsed Process Line %s \n", line)
		if len(line) > 0 {
			results = append(results, parseProcessLine(line))
		}
		if err != nil {
			fmt.Println("Error, Returning Now")
			fmt.Println(err)
			break
		}
	}
	return results
}

func main() {
	processInfoBuf := getProcessInfo()
	processInfo := parseProcessInfo(processInfoBuf)
	for _, p := range processInfo {
		fmt.Printf("Name %s Arguments %s PID %d", p.name, p.arguments, p.pid)
	}
}
