package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgre@123" //psql shell password
	dbname   = "mydatabase"
)

func DataMigration() {
	fmt.Println("setting up....")
	//connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//open database
	db, err = gorm.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	} else {
		fmt.Println("Database Connected Successfully")
	}
	//defer db.close()
	db.AutoMigrate(&User{})

}
