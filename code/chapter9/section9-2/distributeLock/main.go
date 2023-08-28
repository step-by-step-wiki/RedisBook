package main

import (
	"distributeLock/lock"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
	"time"
)

func main() {
	// 创建锁1
	conn1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalf("conn error: %v\n", err)
	}
	distributeLock1 := lock.DistributeLock{
		Key:  "flag",
		TTL:  60,
		Conn: conn1,
	}
	distributeLock1.GenValue()
	defer conn1.Close()

	// 创建锁2
	conn2, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalf("conn error: %v\n", err)
	}
	distributeLock2 := lock.DistributeLock{
		Key:  "flag",
		TTL:  60,
		Conn: conn2,
	}
	distributeLock2.GenValue()
	defer conn2.Close()

	// 创建锁3
	conn3, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalf("conn error: %v\n", err)
	}
	distributeLock3 := lock.DistributeLock{
		Key:  "flag",
		TTL:  60,
		Conn: conn3,
	}
	distributeLock3.GenValue()
	defer conn3.Close()

	// 启动3个goroutine
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go retrieveLock(1, wg, distributeLock1)
	go retrieveLock(2, wg, distributeLock2)
	go retrieveLock(3, wg, distributeLock3)
	wg.Wait()
}

func retrieveLock(i int, wg *sync.WaitGroup, distributeLock lock.DistributeLock) {
	haveRetrieve := false

	for !haveRetrieve {
		// 获取锁
		if distributeLock.Acquire() {
			fmt.Printf("goroutine %d acquire lock success\n", i)
			haveRetrieve = true
		} else {
			fmt.Printf("goroutine %d acquire lock failed, retry soon\n", i)
			time.Sleep(2 * time.Second)
			continue
		}

		// 以等待5s作为假定的业务处理时间
		fmt.Printf("goroutine %d is processing\n", i)
		time.Sleep(5 * time.Second)

		// 释放锁
		if distributeLock.Release() {
			fmt.Printf("goroutine %d release lock success\n", i)
		} else {
			fmt.Printf("goroutine %d release lock failed\n", i)
		}
	}

	wg.Done()
}
