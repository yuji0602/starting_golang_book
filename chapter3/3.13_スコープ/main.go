package main

import (
	"fmt"
	"./foo"
)

func main() {
	/* パッケージのスコープ */
	fmt.Println("パッケージのスコープ")
	fmt.Printf("foo.MAX = %v\n", foo.MAX)
	//fmt.Printf("foo.internal_const = %v\n", foo.internal_const)
	fmt.Printf("foo.FooFunc(5) = %v\n", foo.FooFunc(5))
	//fmt.Printf("foo.internalFunc(5) = %v\n", foo.internalFunc(5))

	/* ファイルのスコープ */
	// $go run main.go app.go か go build -o mainで実行する必要あり
	fmt.Println("ファイルのスコープ")
	fmt.Println("AppVersion:", AppVersion)
	printMessage("Hello")

}
