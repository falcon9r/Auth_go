package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main(){
	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			c, err := r.Cookie("username")
			if err != nil {
				http.ServeFile(w, r , "static/index.html")
			}else{
				fmt.Println(c.Value)
			}
			/*expiration := time.Now().Add(365 * 24 * time.Hour)
			var c = http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
			http.SetCookie(w , &c)
			http.ServeFile(w, r , "static/index.html")	*/
		}
		http.Error(w , "No such path" , http.StatusNotFound)
		return
	})
	http.HandleFunc("/register" , func(w http.ResponseWriter, r *http.Request) {
			var email string = r.PostFormValue("email")
			var password string = r.PostFormValue("password")
			expiration := time.Now().Add(365 * 24 * time.Hour)
			
			value, _ := json.Marshal(fmt.Sprintf(`{"email": %s , "password": %s}`, email , password))
			fmt.Println(string(value))
			c := http.Cookie{Name: "data", Value: string(value), Expires: expiration}
			http.SetCookie(w , &c)	
	})
	println("run...")
	
	server(4000)
}

func server(port int) uint {
	var result string = ":"
	result += strconv.Itoa(port)
	fmt.Print(result)
	if err := http.ListenAndServe(result , nil); err != nil{
		return server(port + 1)
	}
	return 1
}