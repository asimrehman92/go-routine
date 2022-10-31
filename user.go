package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"uid"` //auto generate id
	Name       string       `json:"uname"` // basically json k liay EmpName ki field ka name empname hoga for json
	Location   string       `json:"ulocation"`
}
