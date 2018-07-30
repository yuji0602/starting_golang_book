package main

import (
	"fmt"
	"runtime"
)

func main() {
	go sub()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("無名関数のloop:", i)
		}
	}()
	for i := 0; i < 1000; i++ {
		fmt.Println("main loop: ", i)
	}

	go fmt.Println("Yeah!!")
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())
}

func sub() {
	for i := 0; i < 1000; i++ {
		fmt.Println("sub loop: ", i)
	}
}
