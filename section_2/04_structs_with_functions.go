package main // every file shoud have package declaration

import "log"

type myStruct struct{
    FirstName string
}

// Receiver function tied to myStruct type. This is the way to associate
// f-ns with types. Similar as class methods.
func (m *myStruct) getFirstName() string{
    return m.FirstName
}

func (m *myStruct) setFirstName(name string){
    m.FirstName=name
}

func main()  {

    var myVar myStruct
    myVar.FirstName = "James"
    
    myVar2 := myStruct{
        FirstName: "Mary",
    }

//    log.Println("myVar is set to", myVar.FirstName)
//    log.Println("myVar2 is set to", myVar2.FirstName)
    log.Println("myVar is set to", myVar.getFirstName())
    log.Println("myVar2 is set to", myVar2.getFirstName())

    myVar.setFirstName("Darko")
    log.Println("myVar is set to", myVar.getFirstName())
}
