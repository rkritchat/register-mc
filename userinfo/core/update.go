package core

import "errors"

type UpdateUserInfoReq struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age uint8 `json:"age"`
}

type UpdateUserInfoRes struct {
	Status string `json:"status"`
	Desc string  `json:"desc"`
}

type UpdateService struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age uint8 `json:"age"`
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

func (s UpdateService) ValidateBusinessRule() error {
	if len(s.FirstName) > 10 {
		return errors.New("firstName cannot more than ten")
	}else if len(s.LastName) > 10 {
		return errors.New("lastName cannot more than ten")
	}
	return nil
}

func (s UpdateService) Execute() (string, error) {
	return "Update user detail successfully", nil
}
