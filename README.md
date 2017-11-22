# Winston Chatbot
Winston is a web-app chatbot created using GoLang.

## About
Name: Daire Ní Chatháin
ID: G00334757

This is an adaptation of the Eliza AI chat bot coded in [Go](http://golang.org). This bot was created as part of my Data Representation and Querying Module in GMIT. The code here is adapted from [smallsurething.com](https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/) as well as this example provided by [Ian McLoughlin](https://github.com/data-representation/eliza).

## Eliza
Computer scientist, Joseph Weizenbaum created the ELIZA program in the early 1960's. His goal was to demonstrate the superficiality of communication between man and machine. Eliza is based on simple pattern recognition, based on a stimulus-response model. To give the illusion of understanding, Eliza uses pronoun transformation. This involves manipulating an example user input of "I am worried about college" to give a response such as "Tell me more about why you are worried about college."

## Winston
Winston is a polite, butler-esque chat bot who only wants the best for you. Similar to Eliza, he emulates a Rogerian psychotherapist. Winston asks open ended questions, always getting the user to guide the conversation. He also has a set of unique greetings and responses in keeping with his character. Winston uses phrases such as "Tally-ho!", "Old chum" and "Salutations!". 

### Winston in Action
![Alt Text](https://github.com/DaireNiC/WinstonBot/blob/master/winston_example.gif=450x300)


### How it works
Winston is a  goLang web-app. It serves a HTML page as the root resource. Once the .exe is ran it can be accessed from a browser using the local host and port number(8000). THe application consists of a web chat window similar to that found on websites such a WhatsApp or Facebook's Messenger. The input text box uses AJAX to interact with the chatbot. 

Winston operates by using [RegeX](https://golang.org/pkg/regexp/) to recognise key words or phrases from the user input to reproduce a response using those keywords from preprogrammed responses which can be found in the data/responses.dat file. 

#### The program can be broken into the following steps:
          1. The user enters input.
          2. This input is then prepared for processing.
          3. The input is searched for keywords.
          4. Pronouns are swapped.
          5. The response is displayed in the chat window.
          
### Code Design
In coding this implementation of Eliza I decided not to re-invent the wheel. I built my web app up using previous exercises completed in the problem sheets assigned for this module. To begin I had a simple web app that served a html page. Using AJAX a user can enter input which is then appended onto a list and rendered in the view. 

When it came to adding in the ELIZA functionality I first used my basic implementation completed in my [03_problem sheet](https://github.com/DaireNiC/03_go_problem_sheet). I then decided to add in the more efficient code available here(https://github.com/data-representation/eliza). I built upon this by adding more regeX patterns and unique responses. His enhances the bot by offering a richer conversation experience along with giving the chatbot a quirky personality. I finished by styling to the application using bootstrap.

### Extras 
Making use of what I have learnt in this module, I decided to incorporate cookies in the web app. This can be viewed when the web app loads. A modal opens with an input box where the user can enter their name, this value is then stored as a cookie. If the same user returns to the page the modal opens with a welcome back message and hides the input box.


## Compilation
[Go](https://golang.org) must be installed to run the code.

### Clone this repo
```bash
git clone https://github.com/DaireNiC/WinstonBot
```
### Navigate to the folder
```bash
cd WinstonBot
```
### Build and the web app:
```go
go build winston.go
```
### Run the exe:
```bash
./winston.exe
```
After running the exe open a browser and type in localhost and the port number (8000)
```bash
e.g 127.0.0.1:8000
```
