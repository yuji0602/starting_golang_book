package main

import "fmt"

func init() {
	fmt.Println("init()が先に表示される")
}

func main() {
	fmt.Println("main()")
}
