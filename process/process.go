package main 

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
	state := cmd.ProcessState
	fmt.Printf(" Pid: %d\n", state.Pid())
	fmt.Printf(" System: %d\n", state.SystemTime())
	fmt.Printf(" User: %d\n", state.UserTime())
}