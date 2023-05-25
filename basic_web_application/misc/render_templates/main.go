package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "home.page.html")
}

// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
    parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

    err := parsedTemplate.Execute(w, nil)
    if err != nil {
        fmt.Println("error parsing template",err)
        return
    }
}

// main is the main app function
func main() {
    
    http.HandleFunc("/", Home) 
    http.HandleFunc("/about", About)

    fmt.Println("Starting application on port", portNumber)

    http.ListenAndServe(portNumber, nil)
        
}
