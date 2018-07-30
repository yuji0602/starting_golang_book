package main

import "fmt"

func main () {
	fmt.Println("panic")

	/* panic時でもdeferは実行される */
	defer fmt.Println("panic時でもdeferは実行される")

	/* panicはこれ以上処理を継続しようがない状態を表す */
	panic("runtime error!")
	fmt.Println("Hello World!")
}