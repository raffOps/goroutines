package main

import (
	"fmt"
	"sync"
	"time"
)

func numeros(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		fmt.Printf("%d ", i)
		//time.Sleep(time.Millisecond * 150)
		channel <- i
	}
	print("")
}

func letras(channel chan rune, numberChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for l := 'a'; l < 'n'; l++ {
		// the letters will be printed only after the 6
		for lock := 0; lock < 6 && l == 'a'; {
			lock = <-numberChan
		}
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
		channel <- l
	}
}

func main1() {

	var wg sync.WaitGroup
	wg.Add(2)

	numberChan := make(chan int, 100)
	letterChan := make(chan rune, 100)
	go numeros(numberChan, &wg)
	go letras(letterChan, numberChan, &wg)

	wg.Wait()
	fmt.Printf("Fim")
}

func main() {
	main1()
}
