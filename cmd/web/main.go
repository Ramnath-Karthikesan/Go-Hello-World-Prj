package main

import (
	"fmt"
	"log"
	"myapp/cmd/pkg/config"
	"myapp/cmd/pkg/handlers"
	"myapp/cmd/pkg/render"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	// fmt.Println("Hello world!")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello World!")
	// 	fmt.Println("Number of bytes written", n)

	// 	if err != nil{
	// 		fmt.Println(err)
	// 	}
	// })
	var app config.AppConfig
	
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/divide", handlers.Divide)

	fmt.Println("Starting application on port", portNumber)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

	// fmt.Println("Starting application on port", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)
}
