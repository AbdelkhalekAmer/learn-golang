package main

import "fmt"

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello!"
}

type spanishBot struct{}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
