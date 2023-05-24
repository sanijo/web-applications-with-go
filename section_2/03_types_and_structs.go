package main // every file shoud have package declaration

import (
	"log"
	"time"
)

// To make type/function/var public(visible outside the package) uppercase first
// letter should be used e.g. User
type User struct{
    FirstName string
    LastName string
    PhoneNumber string
    Age int
    BirthDate time.Time
}

func main()  {
    user := User{
        FirstName: "Trevor",
        LastName: "Sawler",
        PhoneNumber: "+385987695",
    }

    var user2 = User{
        FirstName: "Mark",
    }

    log.Println(
        "user FirstName =",user.FirstName,"user BirthDate =",user.BirthDate)
    log.Println("user2 FirstName =",user2.FirstName)
}

