package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
)

func MakeRouter()  {

	router := mux.NewRouter()
	staticRoutes := []string{"css","js","images"}
	initializeStaticRoutes(router, staticRoutes)

	router.HandleFunc("/", MainPageHandler).Methods("GET")
	router.HandleFunc("/post", MainPageHandlerPOST).Methods("POST")

	http.Handle("/", router)
}


func initializeStaticRoutes( router *mux.Router ,staticRoutes []string){
	for i := 0; i < len(staticRoutes); i++{
		router.PathPrefix("/"+staticRoutes[i]+"/").Handler(
			http.StripPrefix("/"+staticRoutes[i]+"/", http.FileServer(http.Dir("./app/"+staticRoutes[i]+"/"))))
	}
}
func MainPageHandler(w http.ResponseWriter, r *http.Request){
	dat, err := ioutil.ReadFile("./app/index.html")
	if err != nil{
		log.Print(err)
	}
	fmt.Fprint(w, string(dat))
}

func MainPageHandlerPOST(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("success")
}
