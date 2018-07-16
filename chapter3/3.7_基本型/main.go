package main

import (
	"fmt"
	"math"
)

func main() {
	/* 論理値型 */
	// bool型の変数bを定義
	var b bool
	// 変数bにtrueを代入
	b = true
	// 型推論を利用した変数定義も可能
	//b := false
	fmt.Println(b)

	/* 数値型 */
	var (
		n1 int
		n2 int64
	)
	n1 = 1
	// 型が異なるためコンパイルエラーとなる
	//n2 = n1
	// キャストすると代入できる
	n2 = int64(n1)
	fmt.Println(n2)

	// 最大値
	fmt.Printf("uint32 max value = %d\n", math.MaxUint32)

	/* 浮動小数点型 */
	// 最大値
	fmt.Printf("float32 max value = %f\n", math.MaxFloat32)
	fmt.Printf("float64 max value = %f\n", math.MaxFloat64)

	// 最小値
	fmt.Printf("float32 min value = %f\n", math.SmallestNonzeroFloat32)
	fmt.Printf("float64 min value = %f\n", math.MaxFloat64)

	// 型変換
	f64 := 1.0
	f32 := float32(1.0)
	fmt.Printf("float64 = %f\n", f64)
	fmt.Printf("float32 = %f\n", f32)

	zero := 0.0
	// +Inf 正の無限大
	pinf := 1.0 / zero
	// -Inf 負の無限大
	ninf := -1.0 / zero
	// NaN 非数
	nan := zero / zero
	fmt.Printf("正の無限大 = %f\n", pinf)
	fmt.Printf("負の無限大 = %f\n", ninf)
	fmt.Printf("非数 = %f\n", nan)

	// 指数
	fmt.Printf("1.0e2(10の2乗) = %v\n", 1.0e2)
	fmt.Printf("1.0e+2(10の2乗) = %v\n", 1.0e+2)
	fmt.Printf("1.0e-2(10のマイナス2乗) = %v\n", 1.0e-2)

	/* 複素数型 */
	// complex128型の変数cを定義して1.0+3iを代入
	c := 1.0 + 3i
	// 定義済み関数complexを使って複素数型の値を生成
	//c := complex(1.0, 3)
	fmt.Printf("complex(1.0, 3i) = %v\n", c)

	/* rune型*/
	// runeはシングルクォートで囲われたUnicode文字で構成
	r := '松'
	fmt.Printf("%v\n", r)

	/* 文字列型 */
	// 文字列はダブルクォートで囲われる
	s := "Goの文字列"
	fmt.Printf("%v\n", s)

	// RAW文字列リテラル
	raw := `
Goの
RAW文字列リテラルによる
複数行に渡る
文字列
\n <- RAWリテラルでは\nはそのままの値が表示される
\n
`
	fmt.Printf("%v\n", raw)
}
