package main // every file shoud have package declaration

import "fmt"

// Declare variable that is available to everything inside the package
var packageAvailableVariable string 

func main()  {

    fmt.Println("Hello!")
    
    var whatToSay string
    var i int
    
    whatToSay = "Some expression..."
    fmt.Println(whatToSay)
    
    i = 69
    fmt.Println("i = ", i)

//    var returnedString string = saySomething()
    returnedString := saySomething() // set var to type returned by f-n
    fmt.Println("The function saySomething() returned:", returnedString)

    // Multiple return from a single f-n
    first, second := multipleReturn()
    fmt.Println("first return var:",first,"\nsecond return var:",second)

}

func saySomething() string {
    return "some expression"
}

func multipleReturn() (string, int) {
    return "some", 2
}
