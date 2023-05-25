package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/pkg/config"
	"webapp/pkg/handlers"
	"webapp/pkg/render"
)

const portNumber = ":8080"

// main is the main app function
func main() {
    
    var app config.AppConfig

    tc, err := render.CreateTemplateCache()
	if err != nil {
        log.Fatal("cannot create template cache")
	}
   
    // Setting TemplateCache in config so that is cached all the time while app
    // is running
    app.TemplateCache = tc
    app.UseCache = false

    repo := handlers.NewRepo(&app)
    handlers.NewHandlers(repo)
    
    // Give access to app config variable inside render package
    render.NewTemplates(&app)

    http.HandleFunc("/", handlers.Repo.Home) 
    http.HandleFunc("/about", handlers.Repo.About)

    fmt.Println("Starting application on port", portNumber)

    http.ListenAndServe(portNumber, nil)
        
} 
