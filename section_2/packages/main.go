package main

import (
	"fmt"
	"github.com/sanijo/packageexercise/helpers"
)

// Create package with: go mod init github.com/sanijo/packagename
// Then create empty directory with module file e.g helpers/helpers.go

func main()  {
    myVar := helpers.SomeType {
        TypeName: "Xman",
        TypeNumber: 69,
    }
    // Access the fields of SomeType
    fmt.Println("Type Name:", myVar.TypeName)
    fmt.Println("Type Number:", myVar.TypeNumber)
}
