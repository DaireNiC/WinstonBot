//Resources: (1) https://ianmcloughlin.github.io

package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

//GreetingData does this
type GreetingData struct {
	Greeting string
}

// Create Resposne array filled with strings that rep responses
var responses = []string{"I'm not sure what you're trying to say. Could you explain it to me?",
	"How does that make you feel",
	"Why do you say that?"}

// substitute pronouns in response
var substitutions = map[string]string{
	"am":       "are",
	"was":      "were",
	"i":        "you",
	"i'd":      "you would",
	"i've":     "you have",
	"i'll":     "you will",
	"my":       "your",
	"are":      "am",
	"you've":   "I have",
	"you'll":   "I will",
	"your":     "my",
	"yours":    "mine",
	"you":      "me",
	"me":       "you",
	"myself":   "yourself",
	"yourself": "myself",
	"i'm":      "you are",
}

func elizaResponse(input string) string {

	//Regular expression to match on "father"
	reg := regexp.MustCompile("(?i)\\bfather|dad(.*)|papa\\b")
	if reg.MatchString(input) {
		return "Why don’t you tell me more about your father?"
	}
	// match on anger/angry
	reg = regexp.MustCompile("(?i)\\b(anger|angry|hate)\\b")
	if reg.MatchString(input) {
		return "Where you think that anger stems from?"
	}
	//match on lonely\loneliness
	reg = regexp.MustCompile("(?i)\\blonel(.*)\\b")
	if reg.MatchString(input) {
		return "You can talk to me. I'm hear to listen"
	}
	//match on childhood/family/
	reg = regexp.MustCompile("(?i)\\b(too)\\b")
	if reg.MatchString(input) {
		return "That sounds a little overwhelming, try take a step back and think about where that comes from"
	}
	//reg expression to check for "I am and variants e.g I'm "
	regIam := regexp.MustCompile(`\b(?i)I'?\s*a?m(.*)`)
	feeling := regIam.FindStringSubmatch(input)

	//return everything after the regEx match
	if len(feeling) > 1 {
		//split every time there is a white space
		inputWords := strings.Split(feeling[1], " ")
		inputWords = strings.Split(feeling[1], " ")
		//check evrey word for a pronoun & swap on a match
		for i := range inputWords {
			if _, found := substitutions[inputWords[i]]; found {
				inputWords[i] = substitutions[inputWords[i]]
			}
		}
		// reassemble the sentence and adding space between each word
		feeling[1] = strings.Join(inputWords, " ")
		response := "How do you know you are%s?"
		return fmt.Sprintf(response, feeling[len(feeling)-1])
	} else {
		//no match will return a random response
		//random num generator
		rand.Seed(time.Now().UTC().UnixNano())
		//Radnomly pick a response from the response array
		random := (rand.Intn(3))
		//return a random response
		return (responses[random])
	}

}

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	// Get the user input from the request.
	input := r.URL.Query().Get("value")
	// send the user input to elizaResponse to be analysed
	// Trim the user input's end of line characters.
	input = strings.Trim(input, "\r\n")
	output := elizaResponse(strings.ToLower(input))
	// If the user input was quit, then quit.
	// Note that Eliza gets to respond to quit before this happens.
	if strings.Compare(strings.ToLower(strings.TrimSpace(input)), "quit") == 0 {
		output = "Good bye old chum!"
	}
	//write winston response
	fmt.Fprintf(w, "%s", output)

}

func generateGreeting() string {
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
	return (greetings[random])
}

func chatSession(w http.ResponseWriter, r *http.Request) {
	//fs := http.FileServer(http.Dir("static"))
	//greeting:= generateGreeting()
	greet := GreetingData{
		Greeting: generateGreeting(),
	}

	t, err := template.ParseFiles("index.html") //parse the html file homepage.html
	if err != nil {                             // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, greet) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {           // if there is an error
		log.Print("template executing error: ", err) //log it
	}

}

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
	*/
	//http.Handle("/", fs)
	//http.Handle("/", fs)
	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}
