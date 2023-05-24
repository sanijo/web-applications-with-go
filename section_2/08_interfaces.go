package main

import "fmt"

// In Go, interfaces define a set of method signatures that a type must implement 
// to be considered as implementing that interface. Interfaces enable polymorphism, 
// allowing different types to be used interchangeably based on shared behavior.

type Animal interface {
    Says() string
    NumberOfLegs() int
}

type Dog struct {
    Name string
    Breed string
}

type Cat struct {
    Name string
    Color string
}

// receiver f-ns (can also be bounded by value (cannot modify original type))
func (d *Dog) Says() string{
    return "woof"
}

func (d *Dog) NumberOfLegs() int  {
   return 4
} 

func (d *Cat) Says() string{
    return "meow"
}

func (d *Cat) NumberOfLegs() int  {
   return 4
} 

func main()  {
    dog := Dog {
        Name: "Rex",
        Breed: "German Shepherd",
    }
    cat := Cat {
        Name: "Miu",
        Color: "White",
    }

    PrintInfo(&dog)
    PrintInfo(&cat)
}

func PrintInfo(a Animal) {
    fmt.Println("This animal says",a.Says(),"and has",a.NumberOfLegs(),"legs")
}
