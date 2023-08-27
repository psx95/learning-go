// everything in go is within a package
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// the function named "main" defines the entrypoint of a go program
func main() {
	// flag is used to accept typed values from command line arguments
	// this line tells go to accept a String value for "level" from the command line
	// the default value for "level" if nothing is provided will be "CRITICAL"
	// and the last argument is a help message that describes the flag.
	level := flag.String("level", "INFO", "Log level to filter for")
	variant := flag.String("variant", "WEB-APP", "Which mode to run the program in")
	flag.Parse() // tells go to look at the command line arguments and populate the variables

	if strings.EqualFold(*variant, "CLI") {
		fmt.Println("Passed custom argument for variant")
		Analyze(*level)
		os.Exit(0)
	} else {
		fmt.Println("Running webserver at localhost:3000")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Response for request %s", r.URL)
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html")
	})

	http.ListenAndServe(":3000", nil)
}
