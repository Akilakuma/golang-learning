package main

import (
	"fmt"
	"sync"
)

func main() {

	// 失敗
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Error:", err)
	// 	}
	// }()

	var wg sync.WaitGroup
	testA(&wg)
	wg.Wait()
}

func testA(wg *sync.WaitGroup) {

	// 失敗
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Error:", err)
	// 		wg.Done()
	// 	}
	// }()

	for i := 0; i < 10; i++ {
		// 失敗
		// defer func() {
		// 	if err := recover(); err != nil {
		// 		fmt.Println("Error:", err)
		// 		wg.Done()
		// 	}
		// }()

		count := i
		wg.Add(1)
		go testB(wg, count)

	}
}

func testB(wg *sync.WaitGroup, count int) {

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Error:", err)
	// 		wg.Done()
	// 	}
	// }()

	fmt.Println(count)
	if count == 5 {
		panic("occur panic!")
	}

	fmt.Println("doing something")
	wg.Done()
}
