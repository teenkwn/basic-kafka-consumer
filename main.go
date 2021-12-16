package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("okay")
	appKafka.StartKafka()
	fmt.Println("kafka has been started")
	time.Sleep(10 * time.Minute)
}
