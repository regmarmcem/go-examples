package main 

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	fifoName := "exec.fifo"
	if err := unix.Mkfifo(fifoName, 0o622); err != nil {
		panic(err)
	}
 
	fds, err := unix.Socketpair(unix.AF_LOCAL, unix.SOCK_STREAM|unix.SOCK_CLOEXEC, 0); 
	if err != nil {
		panic(err)
	}
	parentInitPipe := os.NewFile(uintptr(fds[1]),"~init-p")
	childInitPipe := os.NewFile(uintptr(fds[0]), "init-c")
	fmt.Printf("pair: %d", parentInitPipe)
	fmt.Printf("pair: %d", childInitPipe)
}