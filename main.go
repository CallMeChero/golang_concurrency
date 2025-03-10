package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) // samo ako znamo koja je najsporija operacija
}

// func main() {
// 	done := make(chan bool)
// 	go greet("Nice to meet you!", done)
// 	go greet("How are you?", done)
// 	go slowGreet("How ... are ... you ...?", done)
// 	go greet("I hope you're liking the course!", done)

// 	// ako koristimo jedan channel za vise goroutines onda moramo cekati za isto toliko vrijednosti koje su definirane u rutinama
// 	<-done
// 	<-done
// 	<-done
// 	<-done
// }

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	// ako koristimo jedan channel za vise goroutines onda moramo cekati za isto toliko vrijednosti koje su definirane u rutinama
	// for _, done := range dones {
	// 	<-done
	// }

	for range done {

	}
}
