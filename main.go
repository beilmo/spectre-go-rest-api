package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/beilmo/spectre-go-rest-api/infrastructure/logging"
	"github.com/beilmo/spectre-go-rest-api/infrastructure/router"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	fmt.Println("Hello World!")
	i := 101
	fmt.Println(i)

	router := router.NewRouter(logging.ConsoleLogger{})
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8088", router))
}
