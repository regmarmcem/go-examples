package main

import (
	"fmt"
)

func main() {
	gosh_loop();
	return EXIT_SUCCESS;
}

func goshLoop() {
	var line String
	var args String
	var status bool
	
	for {
		fmt.Println("> ")
		line = readLine()
		args = splitLine()
		status = execute()
		if (!status) {
			break
		}
	}
}