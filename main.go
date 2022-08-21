package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r , "static/index.html")	
		}
		http.Error(w , "No such path" , http.StatusNotFound)
		return
	})
	println("run...")
	var result int = 0;
	
	if err := http.ListenAndServe(":4040" , nil); err != nil{
		log.Fatal(err)
	}
	fmt.Scan(&result)
	if(result != 0){
		panic("end...")
	}
	
}