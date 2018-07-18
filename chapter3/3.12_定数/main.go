package main

import (
	"fmt"
	"math"
)

const ONE = 1

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}

func OmittedValue() {
	// 値が設定されていない場合はコンパイルエラーとなる
	//const (
	//	X
	//	Y
	//	Z
	//)
	const (
		X  = 1
		Y
		Z
		S1 = "あ"
		S2
	)
	fmt.Printf("X=%d, Y=%d, Z=%d, S1=%s, S2=%s\n", X, Y, Z, S1, S2)
}

func ConstantValueEquation() {
	const (
		X  = 2
		Y  = 7
		Z  = X + Y
		S1 = "今日"
		S2 = "晴れ"
		S  = S1 + "は" + S2
	)
	fmt.Printf("Z=%d, S=%s\n", Z, S)

}

func ConstantType() {
	//const (
	//	I64 int64 = -1
	//	F64 float64 = 1.2
	//)
	const (
		I64 = int64(-1)
		F64 = float64(1.2)
	)
	fmt.Printf("I64=%d, F64=%v\n", I64, F64)
	/* 論理値の定数 */
	const B = 5 < 2
	fmt.Printf("B=%v\n", B)

	/* 浮動小数点の定数 */
	const (
		FL32 = float32(math.Pi)
		FL64 = float64(math.Pi)
	)
	fmt.Printf("FL32=%v\n", FL32)
	fmt.Printf("FL64=%v\n", FL64)

	/* 複素数の定数 */
	const (
		C = 4.7 + 1.3i
	)
	fmt.Printf("C=%v\n", C)

	/* ルーン、文字列の定数 */
	const (
		R = 'あ' // rune
		S = "Go言語"
		RS = `今日は良い天気
明日もきっと良い天気のはず。`
	)
	fmt.Printf("R=%v, S=%v, RS=%v\n", R, S, RS)
}

func Iota() {
	const (
		A = iota
		B = iota
		C = iota
	)
	fmt.Printf("A=%d, B=%d, C=%d\n", A, B, C)
	// 省略した場合も上記と同じ様になる
	//const (
	//	A = iota
	//	B
	//	C
	//)
	//fmt.Printf("A=%d, B=%d, C=%d", A, B, C)
	const (
		AA = 1 + iota
		BB
		CC
	)
	fmt.Printf("AA=%d, BB=%d, CC=%d\n", AA, BB, CC)
}

func main() {
	/* 定数 */
	fmt.Println("定数")
	x, y := one()
	fmt.Printf("x=%d, y=%d\n", x, y)

	/* 値の省略 */
	fmt.Println("値の省略")
	OmittedValue()

	/* 定数値の式 */
	fmt.Println("定数値の式")
	ConstantValueEquation()

	/* 定数の型 */
	fmt.Println("定数の型")
	ConstantType()

	/* iota */
	fmt.Println("iota")
	Iota()

}
