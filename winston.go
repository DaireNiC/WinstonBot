//Resources: 
// https://ianmcloughlin.github.io :: 2017-11-02

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	// Get the user input from the request.
	input := r.URL.Query().Get("value")
	// Declare an output string.
	var output string;

	// Try to convert the user input from a string to an unsigned 64-bit integer.
	if uintinput, err := strconv.ParseUint(input, 10, 64); err == nil {
		// If we got an integer, and it's less than 10 then convert it to a word.
		if uintinput < 10 {
			unit := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
			output = unit[uintinput]
		} else {
			// Otherwise just return the input.
			output = input
		}
	} else {
		output = input
	}
	fmt.Fprintf(w, "%s", output)
}

func main() {

	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}