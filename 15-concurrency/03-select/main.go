package main

import (
	"fmt"
	"time"
)

func main() {
	rabi := make(chan string)
	kpOli := make(chan string)
	prachanda := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		rabi <- "New PM is Rabi"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		kpOli <- "New PM is KP Oli"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		prachanda <- "New PM is Prachanda"
	}()

	// fmt.Println(<-kpOli)
	// fmt.Println(<-prachanda)
	// fmt.Println(<-rabi)

	select {
	case message := <-rabi:
		fmt.Println(message)
	case message := <-kpOli:
		fmt.Println(message)
	case message := <-prachanda:
		fmt.Println(message)
	}
}
