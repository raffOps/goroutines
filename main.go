package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(numberChan chan<- int, letterChan <-chan rune, wg *sync.WaitGroup) {
	defer close(numberChan)
	defer wg.Done()

	for i := 0; i < 20; i++ {
		var letter rune
		var ok bool
		for i == 8 && letter <= 'g' { // after 8, the next number will be printed only after the letter g
			letter, ok = <-letterChan
			if !ok {
				break
			}
		}
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
		numberChan <- i
	}
}

func printLetters(letterChan chan<- rune, numberChan <-chan int, wg *sync.WaitGroup) {
	defer close(letterChan)
	defer wg.Done()
	for l := 'a'; l <= 'm'; l++ {
		var number int
		var ok bool
		for l == 'a' && number <= 6 { // the letters will be printed only after the 6
			number, ok = <-numberChan
			if !ok {
				break
			}
		}
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
		letterChan <- l
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	numberChan := make(chan int, 100)
	letterChan := make(chan rune, 100)
	go printNumbers(numberChan, letterChan, &wg)
	go printLetters(letterChan, numberChan, &wg)

	wg.Wait()
	fmt.Printf("End")
}
