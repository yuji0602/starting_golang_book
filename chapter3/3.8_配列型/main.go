package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", a[0])
	// インデックスが範囲を超えるとエラー
	//fmt.Printf("%v\n", a[5])
	// インデックスに負の数を指定するとエラー
	//fmt.Printf("%v\n", a[-1])

	// 初期値なしの場合、[0, 0, 0, 0, 0]が設定される
	ar := [5]int{}
	fmt.Printf("%v\n", ar)

	// 様々な型の初期値
	ia := [3]int{}
	ua := [3]uint{}
	ba := [3]bool{}
	fa := [3]float64{}
	ca := [3]complex128{}
	ra := [3]rune{}
	sa := [3]string{}
	fmt.Printf("%v\n", ia)
	fmt.Printf("%v\n", ua)
	fmt.Printf("%v\n", ba)
	fmt.Printf("%v\n", fa)
	fmt.Printf("%v\n", ca)
	fmt.Printf("%v\n", ra)
	fmt.Printf("%v\n", sa)

	/* 要素数の省略 */
	a1 := [...]int{1, 2, 3}
	a2 := [...]int{1, 2, 3, 4, 5}
	a3 := [...]int{}
	fmt.Println("要素数の省略")
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", a3)

	/* 要素への代入 */
	aa := [...]int8{1, 2, 3}
	aa[0] = 0
	aa[2] = 0
	fmt.Println("要素への代入")
	fmt.Printf("%v\n", aa)
	// 要素の型と互換性のない整数の代入はエラーになる
	//aa[2] = 256
	//fmt.Printf("%v\n", aa)

	/* 配列型の互換性 */
	fmt.Println("配列型の互換性")
	//var (
	//	ar1 [3]int
	//	ar2 [5]int
	//)
	//// 要素数が異なるためエラー
	//ar1 = ar2
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}
	// 代入
	arr1 = arr2
	fmt.Printf("%v\n", arr1)
	arr1[0] = 0
	arr1[2] = 0
	fmt.Printf("%v\n", arr1)
}
