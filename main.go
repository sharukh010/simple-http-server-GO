package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter,r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"ParseForm() err: %s",err)
		return 
	}
	fmt.Fprintf(w,"POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Address = %s\n",address)
}
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
	http.HandleFunc("/form",formHandler)
	
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081",nil);err != nil{
		log.Fatal(err)
	}

}