package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the homepage")
}

// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
    sum := addValues(2, 3)
    message :=  fmt.Sprintf("This is about page and 2 + 3 = %d", sum)
    fmt.Fprintf(w, message)
}

// Divide is divide page handler
func Divide(w http.ResponseWriter, r *http.Request) {
    var num1 float32 = 2.0
    var num2 float32 = 0.0
    f, err := divideValues(num1, num2)
    if err != nil {
        fmt.Fprintf(w, "Cannot divide by 0")
        return
    }
    message :=  fmt.Sprintf("%f divided by %f is equal to %f", num1, num2, f)
    fmt.Fprintf(w, message)
}

// addValues adds two int and returns sum
func addValues(x, y int) int {
    return x + y
}

// divideValues adds two int and returns sum
func divideValues(x, y float32) (float32, error) {
    if y <= 0 {
        err := errors.New("cannot divide by 0")
        return 0, err
    }
    result := x/y
    return result, nil
}

// main is the main app function
func main() {
    
    http.HandleFunc("/", Home) 
    http.HandleFunc("/about", About)
    http.HandleFunc("/divide", Divide)

    fmt.Println("Starting application on port", portNumber)

    http.ListenAndServe(portNumber, nil)
        
}
