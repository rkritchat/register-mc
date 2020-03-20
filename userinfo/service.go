package userinfo

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"register/abs"
	"register/userinfo/core"
)

type UpdateUserInfoRes struct {
	Status string `json:"status"`
	Desc string  `json:"desc"`
}

func initDatabase() sql.DB {
	db, err := sql.Open("mysql", "root:P@ssw0rd@tcp(127.0.0.1:33060)/userinfo")
	if err!= nil {
		panic("Cannot connect database")
	}
	return *db
}

func RegisterUserInfo(w http.ResponseWriter, r *http.Request){
	var rj core.RegisterReq
	json.NewDecoder(r.Body).Decode(&rj)
	rs, err := execute(core.RegisterService{RegisterReq: core.RegisterReq{FirstName: rj.FirstName, LastName: rj.LastName, Age: rj.Age}})
	response := initResponse(rs, err)
	generateResponse(w, response)
}

func UpdateUserInfo(w http.ResponseWriter, r *http.Request){
	var rj core.UpdateUserInfoReq
	json.NewDecoder(r.Body).Decode(&rj)
	rs, err := execute(core.UpdateService{UpdateUserInfoReq: core.UpdateUserInfoReq{Id:rj.Id, FirstName: rj.FirstName, LastName: rj.LastName, Age: rj.Age}})
	response := initResponse(rs, err)
	generateResponse(w, response)
}

func execute(i abs.ServiceInterface) (string, error){
	database := initDatabase()
	defer database.Close()
	if err := i.ValidateRequestMsg(); err != nil {
		return "", err
	}
	if err := i.ValidateBusinessRule(&database); err != nil {
		return "", err
	}
	return i.Execute(&database)
}

func initResponse(rs string, err error) UpdateUserInfoRes {
	if err != nil {
		return UpdateUserInfoRes{Status: "Failed", Desc: err.Error()}
	}
	return UpdateUserInfoRes{Status: "Success", Desc: rs}
}

func generateResponse(w http.ResponseWriter, r UpdateUserInfoRes) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(r)
}