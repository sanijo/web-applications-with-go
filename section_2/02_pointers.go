package main // every file shoud have package declaration

import "fmt"

func main()  {

    // first way
    var myString string
    myString = "green"
    // second way
    //myString := "green"
    fmt.Println("myString is set to:", myString)

    changeUsingPointer(&myString)
    fmt.Println("myString is modified to:", myString)

}

func changeUsingPointer(s *string){
    fmt.Println("adress of s is:", s)
    newValue := "red"
    *s = newValue
}

