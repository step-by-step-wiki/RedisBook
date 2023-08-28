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

	psc := redis.PubSubConn{Conn: conn}
	// 订阅 MQChannel 频道
	err = psc.Subscribe("MQChannel")
	if err != nil {
		log.Fatal(err)
	}

	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			log.Printf("Received message from channel %s: %s\n", v.Channel, v.Data)
			if string(v.Data) == "exit" {
				err = psc.Unsubscribe("MQChannel")
				if err != nil {
					log.Fatal(err)
				}
				return
			}

		case redis.Subscription:
			log.Printf("%s: %s %d\n", v.Kind, v.Channel, v.Count)
			// 当取消订阅后 退出循环
			if v.Count == 0 {
				return
			}

		case error:
			log.Fatal(v)
		}
	}
}
