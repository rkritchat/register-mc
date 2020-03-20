package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"register/userinfo"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/update", userinfo.UpdateUserInfo).Methods("POST")
	r.HandleFunc("/register", userinfo.RegisterUserInfo).Methods("POST")
	http.ListenAndServe(":8080",r)
}