package core

import (
	"database/sql"
	"errors"
)

type UpdateUserInfoReq struct{
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age uint8 `json:"age"`
}

type UpdateService struct {
	UpdateUserInfoReq
}

func (s UpdateService) ValidateRequestMsg() error {
	if len(s.FirstName) == 0 {
		return errors.New("firstName is required")
	}else if len(s.LastName) == 0{
		return errors.New("lastName is required")
	}else if s.Age <= 0 {
		return errors.New("age must more than zero")
	}
	return nil
}

func (s UpdateService) ValidateBusinessRule(db *sql.DB) error {
	if s.Id < 0 {
		return errors.New("id cannot less than zero")
	}else if len(s.FirstName) > 10 {
		return errors.New("firstName cannot more than ten")
	}else if len(s.LastName) > 10 {
		return errors.New("lastName cannot more than ten")
	}

	rs, err := db.Query("SELECT * FROM user_info WHERE id = ?", s.Id)
	if err != nil {
		panic(err.Error())
	}
	if !rs.Next() {
		return errors.New("invalid id")
	}
	return nil
}

func (s UpdateService) Execute(db *sql.DB) (string, error) {
	upDB, err := db.Prepare("UPDATE user_info SET first_name = ?, last_name = ?, age = ? WHERE id = ?")
	if err != nil{
		panic(err.Error())
	}
	upDB.Exec(s.FirstName, s.LastName, s.Age, s.Id)
	return "Update user detail successfully", nil
}
