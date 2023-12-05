package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	amqp10 "github.com/Azure/go-amqp"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

func amqp091_conn() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
	} else {
		defer conn.Close()
	}

}
func amqp10_conn() {
	conn, err := amqp10.Dial(context.TODO(), "amqp://localhost", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		defer conn.Close()
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [amqp091|amqp10] <number of connections>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid number of connections: ", err)
	}

	var wg sync.WaitGroup
	for i := 1; i <= n; i++ {
		wg.Add(1)
		fmt.Println("Starting connection ", i)
		go func() {
			defer wg.Done()
			if os.Args[1] == "amqp091" {
				amqp091_conn()
			} else if os.Args[1] == "amqp10" {
				amqp10_conn()
			} else {
				fmt.Println("Invalid connection type: use amqp091 or amqp10")
				os.Exit(1)
			}
		}()
	}
	wg.Wait()
	time.Sleep(10 * time.Second)
}
