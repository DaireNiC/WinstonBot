//Resources: (1) https://ianmcloughlin.github.io

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//GreetingData does this
type GreetingData struct {
	Greeting string
}

// A data structure representing a term that should be replaced in a string.
// original is a regular expression to be matched, and substitute is a string to replace the match with.
// An example use is to replace the word you with the word me.
type substitution struct {
	original   *regexp.Regexp
	substitute string
}

// A data structure representing a user input and a list of responses to it that winston can give.
// question is a regular expression representing the user input.
// answers is an array of strings, any of which is a reasonable response to question.
// question can capture groups of characters, and elements of answers can use them.
// $1 is the first match, $2 the second, etc.
type response struct {
	question *regexp.Regexp
	answers  []string
}

// winston is a data structure representing a psychoanalyst.
// responses is an array containing elements of type response, as above.
// Likewise, substitutions is an array containing elements of type substitution.
// The order of the elements in both arrays matters - the responses and substitutions are matched in order.
type Winston struct {
	responses     []response
	substitutions []substitution
}

// Method to read in a text file containing substitutions data.
// It takes a single argument, a string, which is the path to the substitutions data file.
// The file should have the following format:
//   All lines that begin with a hash symbol are comments, and are ignored.
//   Each section of the file should begin with at least one blank line.
//   The next line should be a regular expression for what to substitute.
//   The next line should be the new text for the substitution.
//   After that, there should be at least one blank.
// An example substitutions file is given in data/substitutions.txt.

func (me *Winston) readsubstitutions(path string) {

	// Open the file, logging a fatal error if it fails.
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// Set up a line-by-line scanner for the file.
	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanLines)

	// Read through the file line by line.
	// readoriginal is false if we have not yet read the regular expression to match.
	// It is true if we have read the regular expression, and are now looking for the substitution string.
	for readoriginal := false; scanner.Scan(); {
		// Get the text on the current line.
		s := scanner.Text()

		// Decide what to do with the line.
		switch {
		// If the line is blank or starts with a # character then skip it.
		case strings.HasPrefix(s, "#") || len(s) == 0:
			// Do nothing

		// If we haven't read the original, then append an element to the substitutions array.
		// The regualr expression is compiled, and the substitution is left blank for now.
		case readoriginal == false:
			me.substitutions = append(me.substitutions, substitution{original: regexp.MustCompile(s)})
			readoriginal = true
		// Otherwise read the substitution and assign it to the last element of the substitutions array.
		default:
			me.substitutions[len(me.substitutions)-1].substitute = s
			readoriginal = false
		}
	}
}

// Function to read in a text file containing responses data.
// The file should have the following format:
// All lines that begin with a hash symbol are comments, and are ignored.
// This file should have the following format:
//   Each section of the file should begin with at least one blank line.
//   The next line should be a regular expression, matching a user input.
//   Each subsequent line, until a blank line, should contain a response to
//   the usr input. One of these will be chosen at random upon user input.
//   After the responses, there should be at least one blank.
// An example responses file is given in data/responses.txt.
func (me *Winston) readresponses(path string) {
	// Open the file, and quit on an error.
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// Set up a buffer to read the file, line by line.
	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanLines)

	// Loop through the lines of the file, initialising a flag called newsection to true.
	for newsection := true; scanner.Scan(); {
		// Get the next line of the file, assign it to s.
		s := scanner.Text()

		// Decide what to do, based on the following rules.
		// Note that without a condition, switch in Go if just like if-else.
		// Also, the clauses break automatically.
		switch {
		// Do nothing if the line is a comment (begins with #).
		case strings.HasPrefix(s, "#"):
		// If the line is blank, presume we are starting a new section.
		case len(s) == 0:
			newsection = true
		// If newsection is true then create a new response item with the line as a question.
		// Then set newsection to false.
		case newsection == true:
			me.responses = append(me.responses, response{question: regexp.MustCompile(s)})
			newsection = false
		// Otherwise we're just reading a possible response, adding it to the last response item.
		default:
			me.responses[len(me.responses)-1].answers = append(me.responses[len(me.responses)-1].answers, s)
		}
	}
}

// This function accepts a user input, and gives a response as winston.
func (me *Winston) analyse(userinput string) string {
	// Loop through the responses, looking for a match for the user input.
	for _, response := range me.responses {
		if matches := response.question.FindStringSubmatch(userinput); matches != nil {

			// Select a random answer.
			answer := response.answers[rand.Intn(len(response.answers))]

			// Fill the answer with the captured groups from the matches.
			for i, match := range matches[1:] {
				// Reflect the pronouns in the captured group.
				for _, sub := range me.substitutions {
					match = sub.original.ReplaceAllString(match, sub.substitute)
					// Remove any spaces at the start or end.
					match = strings.TrimSpace(match)
				}
				// Replace €1 with the first reflected captured group, €2 with the second, etc.
				answer = strings.Replace(answer, "$"+strconv.Itoa(i+1), match, -1)
			}

			// Clear any ~~ markers from the string. They prevent future matches.
			answer = strings.Replace(answer, "~~", "", -1)

			// Send the filled answer back.
			return answer
		}
	}

	return "I don't know what to say."
}

// Create a new instance of winston.
var winston = Winston{}

func userinputhandler(w http.ResponseWriter, r *http.Request) {

	// Read the substitutions file.
	winston.readsubstitutions("data/substitutions.txt")
	// Read the responses file.
	winston.readresponses("data/responses.txt")
	// Get the user input from the request.
	input := r.URL.Query().Get("value")
	// send the user input to winstonResponse to be analysed
	// Trim the user input's end of line characters.
	input = strings.Trim(input, "\r\n")
	output := winston.analyse(strings.ToLower(input))
	// Note that winston gets to respond to quit before this happens.
	if strings.Compare(strings.ToLower(strings.TrimSpace(input)), "quit") == 0 {
		output = "Good bye old chum!"
	}
	//write winston response
	fmt.Fprintf(w, "%s", output)

}

func generateGreeting(name string) string {
	// Winston greetings
	var greetings = []string{
		"Greetings my friend!",
		"Salutations",
		"Ahoy hoy, How goes you?",
		"How are you today old chap?",
		"Any news old chum?",
	}

	random := (rand.Intn(5))
	//return a random response
	return (greetings[random] + name)
}

func chatSession(w http.ResponseWriter, r *http.Request) {

	http.FileServer(http.Dir("./static"))

	name := ""
	//check if there is a cookie if yes get the name
	var cookie, err = r.Cookie("username")
	if err == nil {
		// If we could read it, try to convert its value to an int.
		name = cookie.Value
	}
	// if not get the name from input
	// Get the user name from the request.
	//name := r.URL.Query().Get("name")

	// Create a cookie instance and set the cookie.
	// You can delete the Expires line (and the time import) to make a session cookie.
	cookie = &http.Cookie{
		Name:    "username",
		Value:   r.URL.Query().Get("name"),
		Expires: time.Now().Add(72 * time.Hour),
	}
	http.SetCookie(w, cookie)

	//	fmt.Fprintf(w, "%s", generateGreeting((name)))

	fmt.Fprintf(w, "%s", name)

}

// Program entry point.
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/chat-session", chatSession)

	//http.HandleFunc("/", chatSession)
	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8000", nil)

}

/*
func main() {

	//random num generator
	rand.Seed(time.Now().UTC().UnixNano())
	//Radnomly pick a response from the greetings array
	//return a random response
	// parse the session html file

	http.HandleFunc("/", chatSession)
	//	t, _ = template.ParseFiles("/static/index.html")
	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go

	//fs := http.FileServer(http.Dir("static"))
	//http.Handle("/", fs)

	//tmpl, _ := template.ParseFiles("/static/session.html")
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.FileServer(http.Dir("static"))
			greet := GreetingData{
				Greeting: "hello",
			}

			tmpl.Execute(w, greet)

		})

	//http.Handle("/", fs)
	//http.Handle("/", fs)
	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}
*/
