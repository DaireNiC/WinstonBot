# Winston Chatbot
Winston is a web-app chatbot created using GoLang.

## About
Name: Daire Ní Chatháin
ID: G00334757

This is an adaptation of the Eliza AI chat bot coded in [Go](http://golang.org). This bot was created as part of my Data Represenation and Querying Module in GMIT. The code here is adapted from [smallsurething.com](https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/) and also this example provided by [Ian McLoughlin](https://github.com/data-representation/eliza).

## Eliza
Computer scientist, Joseph Weizenbaum created the ELIZA program in the early 1960's. His goal was to demonstrate the superficiality of communication between man and machine. Eliza is based on simple pattern recognition, based on a stimulus-response model. To give the illusion of understanding, Eliza uses pronoun transformation. This involves manipulating an example user input of "I am worried about college" to give a response such as "Tell me more about why you are worried about college."

## Winston
Winston is a polite, butler-esque chat bot who only want the best for you. Similar to Eliza, he emulates a Rogerian psychotherapist. Winston asks open ended questions, always getting the user to guide the conversation. He also has a set of unique greetings and responses in keeping with his character. Winston uses phrases such as "Tally-ho!", "Old chum" and "Salutations!". 

### How it works
Winston is a  goLang web-app. It serves a HTML page as the root resource. Once the .exe is ran it can be accesed from a borwser using the local host and port number(8000). THe application consists of a web chat window similar to that found on websites such a WhatsApp or Facebook's Messenger. Here all interaction between the user and Winston can bee seen.

Winston operates by using [RegeX](https://golang.org/pkg/regexp/) to recognise key words or phrases from the user input to reproduce a response using those keywords from pre–programmed responses which can be found in the data/responses.dat file. 

 Markup : 1. The user enters input.
          2. This input is then prepared for processing.
          3. The input is searched for keywords.
          4. Pronouns are swapped.
          5. The response is displayed in the chat window.



## Compilation
[Go](https://golang.org) must be installed to run the code.

###Clone this repo
```bash
git clone https://github.com/DaireNiC/WinstonBot
```
### Navigate to the folder

```bash
cd WinstonBot
```
### Build and the web app:

```go
go build ./winston.go
```
### Run the exe:
```bash
./winston.exe
```
After running the exe open a borwser and type in localhost and the port number (8000)
```bash
e.g 127.0.0.1:8000
```
