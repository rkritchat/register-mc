package core

import (
	"database/sql"
	"errors"
)

type RegisterReq struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age uint8 `json:"age"`
}

type RegisterService struct {
	RegisterReq
}

func (s RegisterService) ValidateRequestMsg() error {
	if len(s.FirstName) == 0 {
		return errors.New("firstName is required")
	}else if len(s.LastName) == 0{
		return errors.New("lastName is required")
	}else if s.Age <= 0 {
		return errors.New("age must more than zero")
	}
	return nil
}

func (s RegisterService) ValidateBusinessRule(db *sql.DB) error {
	if len(s.FirstName) > 10 {
		return errors.New("firstName cannot more than ten")
	}else if len(s.LastName) > 10 {
		return errors.New("lastName cannot more than ten")
	}
	sel, err := db.Query("SELECT * FROM user_info WHERE first_name = ?", s.FirstName)
	defer sel.Close()
	if err != nil{
		panic(err.Error())
	}
	if sel.Next() {
		return errors.New("duplicate first name")
	}
	return nil
}

func (s RegisterService) Execute(db *sql.DB) (string,error){
	insert, err := db.Prepare("INSERT INTO userinfo.user_info (first_name, last_name, age) VALUES (?,?,?)")
	defer insert.Close()
	if err != nil{
		panic(err.Error())
	}
	insert.Exec(s.FirstName, s.LastName, s.Age)
	return "Register successfully", nil
}
