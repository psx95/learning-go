package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Analyze(level string) {
	fmt.Printf("Looking for log lines with level %s", level)

	// ACQUIRING A RESOURCE
	logFile, err := os.Open("./log.txt")

	// HANDLE ERRORS DURING RESOURCE ACQUISITION
	if err != nil {
		// Failed to open log file
		log.Fatal(err) // crash the application
	}

	// RELEASE THE RESOURCE ONCE ITS CONSUMEED (DEFERRED)
	// logFile is a handle to a resource on the system, so it needs to be closed in a timely manner
	// defers the execution of the close function until after the main function exits
	defer logFile.Close()

	bufferedReader := bufio.NewReader(logFile)
	// the expression after the semicolon is the loop invariant
	// So keep looping till the err is nil
	// the last part of the for loop statement is the loop increment / conditional update
	for line, err := bufferedReader.ReadString('\n'); err == nil; line, err = bufferedReader.ReadString('\n') {
		if strings.Contains(line, level) {
			fmt.Println(line)
		}
	}
}
