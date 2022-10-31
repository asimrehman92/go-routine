package main

// "time"

// func compute(value int) {
// 	for i := 0; i < value; i++ {
// 		time.Sleep(time.Second) //sleep for 1 second before next iteration
// 		fmt.Println(i)
// 	}
// }

// func worker(ID int, jobs <-chan int, results chan<- int) {
// 	for job := range jobs {
// 		fmt.Println("Worker ", ID, " is working on job ", job)
// 		duration := time.Duration(rand.Intn(1e3)) * time.Microsecond
// 		time.Sleep(duration)
// 		fmt.Println("Worker ", ID, " completed work on job ", job, " within ", duration)
// 		results <- ID
// 	}
// }

// func main() {
// 	fmt.Println("Concurrency with GoRoutines")
// go compute(5)
// go compute(5)

// jobs := make(chan int, 10)
// results := make(chan int, 10)

// // 3 workers
// for x := 1; x <= 3; x++ {
// 	go worker(x, jobs, results)
// }

// // givem them jobs
// for j := 1; j <= 6; j++ {
// 	jobs <- j
// }

// close(jobs)

// for r := 1; r <= 6; r++ {
// 	fmt.Println("result received from worker: ", <-results)
// }

// 	fmt.Scanln() //if we use this we can put go before both function calls
// }

import (
	"fmt"
	"time"

	// "github.com/jasonlvhit/gocron"
	"github.com/go-co-op/gocron"
	// "github.com/prprprus/scheduler"
	// "github.com/jinzhu/gorm"
)

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
func task() {
	date := time.Now()
	day := date.Weekday().String()
	fmt.Println("Today is ", day)
}

// func taskWithParams(a int, b string) {
// 	fmt.Println(a, b)
// }

// func test(stop chan bool) {
// 	time.Sleep(20 * time.Second)
// 	fmt.Println("Today is not friday")
// 	close(stop)
// }


func main() {

	ch := gocron.NewScheduler(time.Local)

	ch.Every(1).Day().Do(task)
	// ch.Every("M").Seconds().Do(task)
	ch.StartBlocking()

	// getAllUsers()
	// var user User = User{Model: gorm.Model{}, Name: "Koopa", Location: "Pakistan"}
	// db.Create(&user)
	// save in db

	//getting 1st user from table
	// createdAt := users[0].CreatedAt.String()
	// past := users[0].CreatedAt
	//printing creation time
	// fmt.Println(createdAt)

	// var arr []string = strings.Split(createdAt, " ")
	// fmt.Println(arr[1])
	// p := fmt.Println

	// loc, _ := time.LoadLocation("UTC")
	// currentTime := time.Now()
	// p("Current: ", currentTime)

	// then := time.Date(2021, 06, 25, 9, 00, 00, 651387237, time.Local)
	// p("Past Tim:", then)

	// p("created at Hour ", then.Hour()) //hour
	// p("now Hour ", currentTime.Hour())
	// p("created at minute ", then.Minute()) //minute
	// p("now minute ", currentTime.Minute())

	// diffe := currentTime.Sub(then)
	// // p(then.Add(-diff))
	// // abc := time.ParseDuration()
	// p("hours differnce ", int(diffe.Hours()))
	// p("minutes difference", int(diffe.Minutes()))
	// minut := int(diffe.Minutes())
	// if minut > 30 {
	// 	fmt.Println("Hello World")
	// } else {
	// 	p("sorry ")
	// }

}
