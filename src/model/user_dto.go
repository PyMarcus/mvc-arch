package model

import "github.com/PyMarcus/mvc-arch/src/configuration/errs"

type userDomain struct {
	Name     string
	Email    string
	Password string
	Age      int
}

type IUserDomain interface {
	CreateUser() *errs.Errs
	FindUser(string) (*userDomain, *errs.Errs)
}

func NewUserDomain(email, password, name string, age int) IUserDomain {
	return &userDomain{
		Name: name,
		Email: email,
		Password: password,
		Age: age,
	}
}
