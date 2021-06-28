package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	PolicyTest()
	// done := make(chan struct{})

	// go func (){
	// 	time.Sleep(time.Second*3)
	// 	done <- struct{}{}
	// }()

	// <- done
	// fmt.Println("Waiting effect")
}

func withWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 5)
	}()

	wg.Wait()
	fmt.Println("Waiting all wait groups")

}
