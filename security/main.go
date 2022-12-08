package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
)

type User struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `jsaon:"name" validate:"required"`
	DA    int    `json:"da" validate:"required"`
}

func main() {
	v := validator.New()
	toto := User{
		Email: "toto",
		DA:    10,
	}
	err := v.Struct(toto)
	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
	}
}
