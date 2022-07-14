package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	// if cmd is "sleep 1000", execution exit immediately but process keep running

	stdout, _ := cmd.StdoutPipe()
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf("(stdout) %s\n", scanner.Text())
		}
	}()
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	
	// if cmd is "cat test.txt", it finishes in 3 seconds and text is shown through stdoutpipe
	time.Sleep(time.Second * 3)
	
	fmt.Printf(" Process: %d\n", cmd.Process)
}