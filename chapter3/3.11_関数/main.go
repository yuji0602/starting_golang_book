package main

import (
	"fmt"
	"errors"
)

var plusAlias = plus

func main() {
	/* 関数定義の基本 */
	fmt.Println("関数定義の基本")
	fmt.Printf("plus(2, 3) == %v\n", plus(2, 3))

	/* 戻り値のない関数 */
	fmt.Println("戻り値のない関数")
	hello()

	/* 複数の戻り値 */
	fmt.Println("複数の戻り値")
	q, r := div(19, 7)
	fmt.Printf("商=%d 剰余=%d\n", q, r)

	/* 戻り値の破棄 */
	//q, _ := div(19, 7)
	//_, r := div(19, 7)
	//// 全ての戻り値を破棄する
	//_, _ := div(19, 7)
	//// 関数呼び出しのみ書くのと同じ
	//div(19, 7)

	/* 関数とエラー処理 */
	fmt.Println("関数とエラー処理")
	_, err := doSomethingError()
	if err != nil {
		fmt.Printf("エラーが発生しました。 エラー=%s\n", err)
	}

	/* 無名関数 */
	fmt.Println("無名関数")
	f := func(x, y int) int { return x + y }
	fmt.Printf("f(2, 3) == %d\n", f(2, 3))
	fmt.Printf("f()の型は %T\n", f)
	fmt.Printf("%#v\n", func(x, y int) int { return x + y })
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3))

	/* 名前付き関数と無名関数 */
	fmt.Println("名前付き関数と無名関数")
	fmt.Printf("plusAlias(2, 3) == %v\n", plusAlias(2, 3))

	/* 関数を返す関数 */
	fmt.Println("関数を返す関数")
	fn := returnFunc()
	fn()
	returnFunc()()

	/* 関数を引数に取る関数 */
	fmt.Println("関数を引数に取る関数")
	callFunction(func() {
		fmt.Println("I'm a function!")
	})

	/* クロージャとしての無名関数 */
	fmt.Println("クロージャとしての無名関数")
	fc := later()
	fmt.Println(fc("Golang"))
	fmt.Println(fc("is"))
	fmt.Println(fc("awesome!"))

	/* クロージャによるジェネレータの実装 */
	fmt.Println("クロージャによるジェネレータの実装")
	ints := integres()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := integres()
	fmt.Println(otherInts())
	fmt.Println(otherInts())
}

func plus(x int, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello, World!")
	return
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func doSomethingError() (string, error) {
	return "", errors.New("Something Error!")
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function.")
	}
}

func callFunction(f func()) {
	f()
}

func later() func(string) string {
	// 1つ前に与えられた文字絵r津を保存するための変数
	var store string
	// 引数に文字列をとり文字列を返す関数を返す
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integres() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}