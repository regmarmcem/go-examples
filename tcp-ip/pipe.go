package main

import (
	"bufio"
	"fmt"
	"os"
)

func pipe() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
}
