package main

import (
	"errors"
	"fmt"
)
// Following command shows out in brower which functions are or are not covered
// in testing:
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out

func divide(x float32, y float32) (float32, error){
    var result float32

    if y == 0 {
        return result, errors.New("Cannot divide by 0.")
    } 
    
    result = x/y
    return result, nil
}

func main()  {
    result, err := divide(100.0, 0.0)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Result of division is",result)
}
