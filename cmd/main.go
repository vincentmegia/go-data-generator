package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/vincentmegia/go-data-generator/internals/models"
	"github.com/vincentmegia/go-data-generator/internals/services"
)

type Data struct {
	Index int
	Item  models.User
}

func main() {
	var queue = make(chan Data)
	waitGroup := sync.WaitGroup{}

	fmt.Print("Enter amount of data to be generated followed by batch limit: ")
	var total, batchLimit int
	fmt.Scan(&total)
	fmt.Print("Enter batch limit: ")
	fmt.Scan(&batchLimit)
	fmt.Println("Batch limit is: ", batchLimit)

	startTime := time.Now()
	waitGroup.Add(2)
	half := total / 2

	// trigger consumer and producer routines
	go producer(queue, half, &waitGroup, "1")
	go producer(queue, half, &waitGroup, "2")
	go consumer(queue, batchLimit, "1")
	go consumer(queue, batchLimit, "2")
	fmt.Println("Starting data generator")
	waitGroup.Wait()
	elapsed := time.Since(startTime)
	fmt.Println("Total time it took: ", elapsed)
	fmt.Println("Existing gracefully...")
}

func producer(queue chan Data, total int, waitGroup *sync.WaitGroup, producer_id string) {
	defer waitGroup.Done()
	for index := 0; index <= total; index++ {
		epoch := strconv.FormatInt(time.Now().UnixMicro(), 8)
		firstname := "john-" + epoch
		lastname := "doe-" + epoch
		username := "john.doe-" + epoch
		password := "salty-joe-" + epoch
		email := "john.doe-" + epoch + "@mail.com"
		user := models.User{}.Create(firstname, lastname, username, password, nil, email, time.Now())
		data := Data{index, user}
		fmt.Println(fmt.Sprintf("[producer-%s] - sending data to queue userid: %s", producer_id, &user.Id))
		queue <- data
	}
}

func consumer(queue chan Data, batchLimit int, consumer_id string) {
	users := []models.User{}
	for data := range queue {
		users = append(users, data.Item)
		if len(users) == batchLimit {
			fmt.Println(fmt.Sprintf("[consumer-%s] - ARRAY BATCH MAX, trigger batch push...", consumer_id))
			userService := services.UserService{}
			userService.BulkInsert(&users)
			users = []models.User{}
		}
	}
}
