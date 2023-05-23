package main

import "fmt"

type bot interface{
	getGreeting() string;
}

type englishBot struct{}
type banglaBot struct{}

func main(){
	eb := englishBot{};
	bb := banglaBot{};

	printGreeting(eb);
	printGreeting(bb);
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting());
}

func (englishBot) getGreeting() string{
	return "English";
}
func (banglaBot) getGreeting() string{
	return "Bangla";
}
