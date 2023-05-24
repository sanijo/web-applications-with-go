package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    HairColor string `json:"hair_color"`
    HasDog bool `json:"has_dog"`
}

func main()  {
    myJson := 
`[
    {
        "first_name": "Clark",
        "last_name": "Kent",
        "hair_color": "black",
        "has_dog": true
    },
    {
        "first_name": "Bruce",
        "last_name": "Wayne",
        "hair_color": "black",
        "has_dog": false
    }
]`

    // Read JSON to struct
    var unmarshalled []Person

    err := json.Unmarshal([]byte(myJson), &unmarshalled)
    if err != nil {
        fmt.Println("Error unmarshalling JSON",err)
    }
    fmt.Println("Unmarshall:",unmarshalled)

    // Write JSON from a struct
    var mySlice []Person

    var m1 Person
    m1.FirstName = "Spider"
    m1.LastName = "Man"
    m1.HairColor = "brown"
    m1.HasDog = false

    var m2 Person
    m2.FirstName = "Tom"
    m2.LastName = "Superhero"
    m2.HairColor = "blue"
    m2.HasDog = true

    mySlice = append(mySlice, m1)
    mySlice = append(mySlice, m2)

    marshalled, err := json.MarshalIndent(mySlice, "", "    ")
    if err != nil {
        fmt.Println("Error marshalling JSON",err)
    }
    fmt.Println("MarshallIndent:",string(marshalled))
}
