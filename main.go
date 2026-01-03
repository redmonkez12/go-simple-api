package main

import (
	"fe-go/internal/app"
	"fe-go/internal/routes"
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	application, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer application.DB.Close()

	r := routes.SetupRoutes(application)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	application.Logger.Printf("we are running on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		application.Logger.Fatal(err)
	}
}
