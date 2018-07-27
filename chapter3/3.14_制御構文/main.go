package main

import "fmt"

func main() {
	/* if */
	funcIf()

	/* ループ */
	funcLoop()

	/* switch */
	funcSwitch()
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
