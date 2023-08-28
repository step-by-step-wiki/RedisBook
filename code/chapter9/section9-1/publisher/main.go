package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mqMessages := []string{"MQMessage1", "MQMessage2", "MQMessage3", "exit"}

	for _, message := range mqMessages {
		reply, err := conn.Do("PUBLISH", "MQChannel", message)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("reply: %v", reply)
	}
}
