package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter text to capitalize")
	// NewReader is wrapper around stdin which provides additonal functionality
	bufferedReader := bufio.NewReader(os.Stdin)
	// Read string till newline character is encountered
	readString, _ := bufferedReader.ReadString('\n')
	readString = strings.TrimSpace(readString)
	readString = strings.ToUpper(readString)
	fmt.Println(readString + "!")
}
