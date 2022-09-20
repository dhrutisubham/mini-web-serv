package main

import(
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 page not found" , http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w,"method is not supported try another method" , http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"Hello")
}

func formHandler(w http.ResponseWriter , r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w,"form parsing error = %v" , err)
		return
	}
	fmt.Fprintf(w,"POST request ucessful")
	name := r.FormValue("name")
	address := r.FormValue("Address")
	fmt.Fprintf(w,"Name = %s\n" ,name)
	fmt.Fprintf(w,"Address = %s\n", address)
	
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form" , formHandler)
	http.HandleFunc("/hello" , hello)

	fmt.Println("Started server at port 8000")

	if err:= http.ListenAndServe(":8000",nil); err != nil{
		log.Fatal(err)
	}
}