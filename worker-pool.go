package main

import (
	"fmt"
	"time"
)

func getAllUsers() {
	var users []User
	db.Find(&users)

	// loop on all users
	p := fmt.Println
	currentTime := time.Now()
	p("Current: ", currentTime)
	for _, v := range users {
		// fmt.Println("Index: ", i, "user name: ", v.Name, "user value: ", (v.CreatedAt).In(time.Local))
		createdAt := (v.CreatedAt).In(time.Local)
		diffe := currentTime.Sub(createdAt)
		minut := int(diffe.Minutes())
		if minut > 30 {
			fmt.Println("Hello ", v.Name, "you are in ")
		}
		// else {
		// 	p("sorry ")
		// }
	}
}
