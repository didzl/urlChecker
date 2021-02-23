package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFail = errors.New("request fail")


func main() {
	c := make(chan string)
	//just doing main func
	people := [5]string{"nico", "han", "ham", "lee", "bae"}
	for _, person := range people {
		go isGood(person, c)
	}
	for i:= 0; i<len(people); i++ {
		fmt.Println(<-c)
	}
}

func isGood(person string, c chan string) {
	time.Sleep(time.Second *5)
	c <- person + " is good"
}