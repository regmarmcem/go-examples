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
	fmt.Printf(" Process: %d\n", cmd.Process)
}