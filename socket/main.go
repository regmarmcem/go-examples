package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	var cflag *bool = flag.Bool("child", false, "Run as a child process")
	flag.Parse()

	if !*cflag {
		fmt.Println("this is parent")
		if err := parent(); err != nil {
			fmt.Printf("parent failed %v\n", err)
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	} else {
		fmt.Println("this is child")
		child()
		os.Exit(0)
	}
}

func parent() (err error) {
	args := []string{"--child"}
	args = append(args, os.Args[1:]...)

	fds, err := generateSocketPair()
	if err != nil {
		fmt.Println("generateSocketPair failed")
		return err
	}

	parentFd := os.NewFile(uintptr(fds[0]), "parent-fd")
	childFd := os.NewFile(uintptr(fds[1]), "child-fd")

	cmd := exec.Command(os.Args[0], args...)
	cmd.ExtraFiles = append(cmd.ExtraFiles, childFd)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	data := []byte("test")
	r := bytes.NewReader(data)
	// parentFd.Write(data)

	if _, err := io.Copy(parentFd, r); err != nil {
		fmt.Printf("[parent] error in io.Copy %v\n", err)
	}
	var test []byte
	test, err = cmd.Output()
	if err != nil {
		fmt.Printf("output failed: %v\n", err)
		return err
	}
	fmt.Println(string(test))
	buf := make([]byte, 4)
	parentFd.Read(buf)
	s := strings.ToLower(string(buf))
	fmt.Printf("[parent] buf is %s", s)

	return nil
}

func child() {
	childfd := os.NewFile(uintptr(3), "child-fd")
	buf := make([]byte, 4)

	if childfd != nil {
		defer childfd.Close()
		childfd.Read(buf)
		s := strings.ToUpper(string(buf))
		fmt.Printf("[child] buf is %s", s)
		r := bytes.NewReader(buf)
		io.Copy(childfd, r)
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func generateSocketPair() (fds [2]int, err error) {
	fds, err = syscall.Socketpair(syscall.AF_LOCAL, syscall.FD_CLOEXEC, 0)
	return fds, err
}
