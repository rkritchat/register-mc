package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"register/userinfo"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/update", userinfo.UpdateUserInfo).Methods("POST")
	http.ListenAndServe(":8080",r)
}