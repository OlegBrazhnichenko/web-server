package main

import (
	"net/http"
)
//
//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()  // parse arguments, you have to call this by yourself
//	fmt.Println(r.Form)  // print form information in server side
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ""))
//	}
//	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
//}

func main() {
	startServer()
}

func startServer(){
	initDB()
	MakeRouter()
	http.ListenAndServe(":9090", nil)
}