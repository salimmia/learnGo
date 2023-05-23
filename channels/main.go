package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}

	c := make(chan string);

	for _, link := range links{
		go checkLinks(link, c);
		// fmt.Println(<- c);
	}

	// for i := 0; i < len(links); i++{
	// 	go checkLinks(<-c, c);
	// }

	for l:= range c{
		
		go checkLinks(l, c);
	}

	// fmt.Println(<- c);
	// fmt.Println(<- c);
}

func checkLinks(link string, c chan string){
	time.Sleep(5 * time.Second);
	_, err := http.Get(link);

	if err != nil{
		fmt.Println(link, "might be down");
		c <- "Might be link down";
		return;
	}

	fmt.Println(link, "is up");

	c <- "Yes its up";
}