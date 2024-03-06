package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(channel chan int, wg *sync.WaitGroup) {
	defer close(channel)
	defer wg.Done()

	for i := 0; i < 12; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
		channel <- i
	}
}

func printLetters(channel chan rune, numberChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for l := 'a'; l < 'n'; l++ {
		// the letters will be printed only after the 6
		var number int
		var ok bool
		for l == 'a' && number < 6 {
			number, ok = <-numberChan
			if !ok {
				break
			}
		}
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
		channel <- l
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	numberChan := make(chan int, 100)
	letterChan := make(chan rune, 100)
	go printNumbers(numberChan, &wg)
	go printLetters(letterChan, numberChan, &wg)

	wg.Wait()
	fmt.Printf("End")
}
