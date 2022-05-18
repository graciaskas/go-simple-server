package main

import(
	"fmt",
	"log",
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.printf("Server running on http://localhost:8000\n")

	if err := http.listenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}