package main

import "fmt"

func main() {
	fmt.Println("recover")

	/* recoverの使い方 */
	fmt.Println("recoverの使い方")
	funcRecover()

	/* panicとrecoverを合わせた使い方 */
	fmt.Println("panicとrecoverを合わせた使い方")
	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})

}

func funcRecover() {
	defer func() {
		if x := recover(); x != nil {
			/* 変数xはpanicに渡されたinterface{} */
			fmt.Println(x)
		}
	}()
	panic("panic!")
	fmt.Println("Hello World!")
}

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Println("panic: unknown")
			}
		}
	}()
	panic(src)
	return
}
