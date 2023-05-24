package main // every file shoud have package declaration

import (
	"fmt"
	"sort"
)

type User struct{
    FirstName string
    LastName string
}

func main()  {
    // Map should be created using := . In [string]string inside [] is index
    // used to lookup the map, and type after [] is value storead at the index.
    // Important thing is that map is immutable. Also, they are not sorted.
    // Interface can be used if you dont know in advance which type you will
    // need: myMap := make(map[string]interface{})
    myMap := make(map[string]string)
    myMap["dog"]="Dzeki"
    fmt.Println(myMap["dog"])
    myMap["dog"]="Bubi"
    fmt.Println(myMap["dog"])

    myMap2 := make(map[string]int)
    myMap2["first"]=1
    myMap2["second"]=2
    fmt.Println(myMap2["first"])
    fmt.Println(myMap2["second"])

    var user1 = User{
        FirstName: "Marko",
        LastName: "Dimsic",
    } 
    user2 := User{
        FirstName: "Luka",
        LastName: "Perkovic",
    }
    myUserMap := make(map[string]User)
    myUserMap["user1"]=user1
    myUserMap["user2"]=user2
    fmt.Println("user1 FirstName",myUserMap["user1"].FirstName)
    fmt.Println("user1 LastName",myUserMap["user1"].LastName)
    fmt.Println("user2 FirstName",myUserMap["user2"].FirstName)
    fmt.Println("user2 LastName",myUserMap["user2"].LastName)

    // Slice is similar to array string myArray [].
    var mySlice []int
    fmt.Println(mySlice)
    mySlice = append(mySlice,4)
    mySlice = append(mySlice,2)
    mySlice = append(mySlice,3)
    fmt.Println(mySlice)
    
    sort.Ints(mySlice)
    fmt.Println(mySlice)

    var numbers1 = []int {2,4,67}
    fmt.Println(numbers1)
    // Shorthand declaration of a slice
    numbers2 := []int{1,2,3,4}
    fmt.Println(numbers2)
    fmt.Println(numbers2[0:2])

    // 2D slice
    slice2d := [][]int{
        {1,2,3},
        {3,4,2},
        {0,9,6},
    }
    fmt.Println("Whole slice2d",slice2d)
    fmt.Println("slice2d[0]",slice2d[0])
    fmt.Println("slice2d[0][2]",slice2d[0][2])
        
    // 2D slice (user specified type)
    slice2duser := [][]User{
        {user1,user1},
        {user2,user2},
    }
    fmt.Println("Whole slice2duser",slice2duser)
    fmt.Println("slice2duser[0]",slice2duser[0])
    fmt.Println("slice2duser[0][1].FirstName",slice2duser[0][1].FirstName)
}
