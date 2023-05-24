package main

import "fmt"

func main()  {
    var isTrue bool 
    isTrue = true

    if isTrue == true {
        fmt.Println("isTrue is",isTrue)
    } else {
        fmt.Println("isTrue is",isTrue)
    }

    cat := "dog"
    if cat == "cat" {
        fmt.Println("cat is cat")
    } else {
        fmt.Println("cat is not cat")
    }

    number := 99
    if number>0 && number<10 {
        fmt.Println("Number",number,"is in range <0,10>")
    } else {
        fmt.Println("Number",number,"is not in range <0,10>")
    }

    option := "bear"
    switch option {
    case "cat":
        fmt.Println("option is set to cat")
    case "dog":
        fmt.Println("option is set to dog")
    case "fish":
        fmt.Println("option is set to fish")
    default:
        fmt.Println("option is set to",option)
    }

}
