package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w,"method is not supported",http.StatusNotFound)
		return 
	}
	fmt.Fprintf(w,"hello! Sharukh Khan")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/hello",helloHandler)
	
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081",nil);err != nil{
		log.Fatal(err)
	}

}