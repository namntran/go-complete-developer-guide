package main

import "fmt"

// interfaces are not generic types,
// interfaces are 'implicit', we don't manually have to say that our custom type satisfies some interface
// interfaces are a contract to help us manage types, helps you re-use code not avoid logic mistakes in your code

type bot interface {
	getGreeting() string
}
type englishBot struct{} // interfaces - type bot implements englishBot
type spanishBot struct{} // interfaces - type bot implements spanishBot

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (englishBot) getGreeting() string {
	// custom logic for generating an english greeting
	return "Hello Guv"
}
func (spanishBot) getGreeting() string {
	// custom logic or generating a spanish greeting
	return "Hola Amigo"
}
