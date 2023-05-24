package main

import "fmt"

func main()  {
    for i:=0; i<=5; i++ {
        fmt.Println("i =",i)
    }

    animals := []string{"cow","monkey","dog"}
    // Range based for loop. _ is used to tell compiler to ignore that return
    // value. If its needed it can be replaced with i for example.
    for _, animal:=range animals {
        fmt.Println("animal =",animal)
    }

    names := map[string]string {
        "name1" : "Dan",
        "name2" : "John",
        "name3" : "Mira",
    }
    for person, name:=range names {
        fmt.Println("user =",person,"name =",name)
    }

    type User struct{
        FirstName string
        LastName string
        Email string
        Age int
    }
    var users []User
    users = append(users, User{"John", "Smith", "john@smith.com", 43})
    users = append(users, User{"Mike", "Smajser", "mike@smajser.com", 32})
    users = append(users, User{"Ivo", "Galic", "ivo@galic.com", 65})
    users = append(users, User{"Pero", "Somer", "pero@somer.com", 19})

    for _,user := range users{
        fmt.Println(user.FirstName, user.LastName, user.Email, user.Age)
    }

}
