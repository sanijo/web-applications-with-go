package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Channels in Go are a core feature of the language for communication and
// synchronization between goroutines (concurrently executing functions or methods).
// They provide a way for goroutines to send and receive values safely
// and efficiently.

func RandomNumber(n int) int {
    rand.Seed(time.Now().UnixNano())
    value := rand.Intn(n)
    return value
}

const numPool int = 100

// Channel can both send and receive if: intChan chan int
// Receive only if: intChan <-chan int
// Send only if: intChan chan<- int
func CalculateValue(intChan chan<- int) {
    randomNumber := RandomNumber(numPool)
    // Push randomNumber to the channel
    intChan <- randomNumber
}

func main()  {
    // Channels (way to send/receive information)
    intChan := make(chan int)
    // When Im done using, close this channel (not necessary)
    defer close(intChan)

    go CalculateValue(intChan)

    num := <- intChan // receive value from channel
    fmt.Println(num) 
}
