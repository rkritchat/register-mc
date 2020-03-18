package userinfo

import (
	"encoding/json"
	"net/http"
	"register/abs"
	"register/userinfo/core"
)

func UpdateUserInfo(w http.ResponseWriter, r *http.Request){
	var rj core.UpdateUserInfoReq
	json.NewDecoder(r.Body).Decode(&rj)
	rs, err := execute(core.UpdateService{FirstName: rj.FirstName, LastName: rj.LastName, Age: rj.Age})
	response := initResponse(rs, err)
	generateResponse(w, response)
}

func execute(i abs.ServiceInterface) (string, error){
	if err := i.ValidateRequestMsg(); err != nil {
		return "", err
	}
	if err := i.ValidateBusinessRule(); err != nil {
		return "", err
	}
	return i.Execute()
}

func initResponse(rs string, err error) core.UpdateUserInfoRes {
	if err != nil {
		return core.UpdateUserInfoRes{Status: "Failed", Desc: err.Error()}
	}
	return core.UpdateUserInfoRes{Status: "Success", Desc: rs,}
}

func generateResponse(w http.ResponseWriter, r core.UpdateUserInfoRes) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(r)
}