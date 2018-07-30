package main

import (
	"fmt"
	"os"
)

func main() {
	/* if */
	funcIf()

	/* ループ */
	funcLoop()

	/* switch */
	funcSwitch()

	/* 型アサーション */
	funcTypeAssertion()

	/* 型によるswitch */
	funcTypeSwitch()

	/* goto */
	funcGoto()

	/* ラベル付き文 */
	funcLabel()

	/* defer */
	funcDefer()

	/* panicとrecoverについては別ファイルにき切り出し */
	/* goについては別ファイルに切り出し */
	/* initについては別ファイルに切り出し */
}

func funcIf() {
	fmt.Println("if")
	// 基本形
	//if x == 1 {
	//}
	// 簡易文付きif
	//if x, y := 1, 2; x < y {
	//	fmt.Printf("x(%d) is less than y(%d)\n", x, y)
	//}
	if n := plus(3, 5); n%2 == 0 {
		fmt.Printf("n(%d) is even\n", n)
	} else {
		fmt.Printf("n(%d) is odd\n", n)
	}
}

func funcLoop() {
	// break
	fmt.Println("for - break")
	var i = 0
	for {
		i++
		if i == 3 {
			fmt.Printf("i == 3, break%v\n", i)
			break
		}
	}
	fmt.Println("-------")

	// 条件式付き
	fmt.Println("for - 条件式付き")
	i = 0
	for i < 3 {
		fmt.Printf("i == %v\n", i)
		i++
	}
	fmt.Println("-------")

	// 古典的
	fmt.Println("for - 古典的")
	for i := 0; i < 3; i++ {
		fmt.Printf("i = %v\n", i)
	}
	fmt.Println("-------")

	// continue
	fmt.Println("for - continue")
	for i := 0; i < 5; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Printf("i = %v\n", i)
		i++
	}
	fmt.Println("-------")

	// 範囲節(Range Clause)
	fmt.Println("for - 範囲節(Range Clause)")
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	// 文字列のループ[index, rune]
	for i, r := range "ABCDE" {
		fmt.Printf("%d=%v\n", i, r)
	}
}

func funcSwitch() {
	// 式によるswitch
	fmt.Println("式によるswitch")
	var n = 3
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}

	// フォールスルー
	fmt.Println("フォールスルー")
	s := "A"
	switch s {
	case "A":
		s += "B"
		fallthrough
	case "B":
		s += "C"
		fallthrough
	case "C":
		s += "D"
		fallthrough
	default:
		s += "E"
	}
	fmt.Printf("s = %s\n", s)

	// 式を伴うswitch
	fmt.Println("式を伴うswitch")
	n = 4
	switch {
	case n > 0 && n < 3:
		fmt.Println("0 < n < 3")
	case n > 3 && n < 6:
		fmt.Println("3 < n < 6")
	}

}

func plus(x int, y int) int {
	return x + y
}

func funcTypeAssertion() {
	fmt.Println("型アサーション")
	/* interface{}には全ての方を指定できる */
	anything(1)
	anything(3.14)
	anything(4 + 5i)
	anything('海')
	anything("日本語")
	anything([...]int{1, 2, 3, 4, 5})

	fmt.Println("型チェック")
	// x.(T)のような式で型を動的にチェック出来る
	var x interface{} = 3
	i := x.(int)
	fmt.Printf("i = %v\n", i)
	// 下記の場合、panic: interface conversion: interface {} is int, not float64というエラーが出る
	//f := x.(float64)
	//fmt.Printf("f = %v\n", f)

	// 以下のようなやり方だとエラーは発生せず、1番目に値、2番目にtrue/falseが入ってくる
	var xx interface{} = 3.14
	i, isInt := xx.(int)
	fmt.Printf("i = %v, isInt = %v\n", i, isInt)
	f, isFloat64 := xx.(float64)
	fmt.Printf("f = %v, isFloat64 = %v\n", f, isFloat64)
	s, isString := xx.(string)
	fmt.Printf("s = %v, isString = %v\n", s, isString)
	// 1つ目の引数に値が必要ない場合、_で省略して書くことが出来る。
	_, isRune := xx.(rune)
	fmt.Printf("isRune = %v\n", isRune)

	/* 変数xはinterface{}型 */
	fmt.Println("interface{}型と型アサーションを組み合わせた判定")
	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Printf("x is integer : %d\n", i)
	} else if s, isString := x.(string); isString {
		fmt.Println(s)
	} else {
		fmt.Println("unsupported type!")
	}
}

func anything(a interface{}) {
	fmt.Println(a)
}

func funcTypeSwitch() {
	fmt.Println("型によるswitch")
	typeSwitch(true)
	typeSwitch(2)
	typeSwitch("あいう")
	typeSwitch('あ')
}

func typeSwitch(x interface{}) {
	/* 変数xはinterface{}型 */
	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int, uint:
		fmt.Println("integer or unsigned integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("don't know")
	}

	/* 型以外に値が必要な場合、v := x.(type)の形で変数に代入 */
	switch v := x.(type) {
	case bool:
		fmt.Println("bool:", v)
	case int:
		fmt.Println(v * v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Printf("%#v\n", v)
	}
	fmt.Println("-----")
}

func funcGoto() {
	print("goto")
	fmt.Println("ここは表示される")
	goto L
	fmt.Println("ここは表示されない")
L: /* ラベル */
	fmt.Println("gotoでラベル L までジャンプ")
}

func funcLabel() {
	fmt.Println("ラベル付き文")
	fmt.Println("ラベル - break")

	LOOP:
		for {
			for {
				for {
					fmt.Println("開始")
					break LOOP
				}
				fmt.Println("ここは通らない")
			}
			fmt.Println("ここは通らない")
		}
	fmt.Println("完了")

	fmt.Println("ラベル - continue")
	L:
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				if j > 1 {
					continue L
				}
				fmt.Printf("%d * %d = %d\n", i, j, i*j)
			}
			fmt.Println("ここは処理されない！")
		}
}

func funcDefer() {
	fmt.Println("defer")
	/* deferに登録された式は関数の終了時に評価される */
	runDefer()

	/* deferはリソースの開放処理で役に立つ */
	file, err := os.Open("../../README.md")
	if err != nil {
		/* ファイルオープンに失敗したらreturn */
		fmt.Println("ファイルオープンに失敗")
		return
	}
	/* ファイルのクローズ処理を登録 */
	defer file.Close()

	/* deferで複数処理させたい場合は無名関数にする */
	defer func() {
		fmt.Println("A")
		fmt.Println("B")
		fmt.Println("C")
	}()
}

func runDefer() {
	defer fmt.Println("deferに登録された式は関数の終了時に評価される")
	defer fmt.Println("deferを複数使うと後に記載したほうがが先に評価される")
	fmt.Println("done")
}