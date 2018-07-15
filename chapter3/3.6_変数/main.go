package main

/*
変数を列挙しているだけなので go run main.go を実行するとエラーになります。
*/

import (
	_ "fmt"
)


/* aaaはパッケージ変数。さまざまな場所から参照可能のため乱用は避けるべき */
var aaa = 100

func main() {
	// int型の変数nを定義する
	var n int
	// int型の変数x, y, zを定義する
	var x, y, z int
	// int型の変数x, yとstring型の変数nameを定義する
	var (
		x, y int
		name string
	)

	// 変数へ代入
	n = 5
	// 型違いのためコンパイルエラー
	n = "string"

	// 複数の変数へ代入
	var x, y int
	x, y = 1, 2
	// 個数違いのためコンパイルエラー
	x, y = 1, 2, 3

	/* 暗黙的な定義 */
	// int型の変数iを定義して1を代入
	i := 1
	// bool型の変数bを定義して真偽値trueを代入
	b := true
	// float64型の変数fを定義して実数3.14を代入
	f := 3.14
	// string型の変数sを定義して文字列"abc"を代入
	s := "abc"
	// int型の変数nが定義され、値には1が代入される
	n := one()

	// 変数への暗黙的定義を利用した代入は1度しか許されない
	i := 1 // int型の変数iを定義して1を代入
	i := 2 // コンパイルエラー

	// 多重定義によるコンパイルエラー
	var i int // int型の変数iを定義
	var i int // コンパイルエラー

	// 変数への再代入には演算子=を用いる
	var a int // int型の変数aを定義
	a = 1 // 変数aに1を代入
	a = 2 // 変数aに2を代入

	b := 1 // int型の変数bを定義して1を代入
	b = 2 // 変数bに2を代入

	/* varと暗黙的な定義 */
	var a = 1 // int型への変数に1を代入
	a := 1 // 上記だと冗長なのでこういう書き方のほうが良い

	// varで変数定義をまとめる書き方
	var (
		n = 1
		s = "string"
		b = true
	)

	// 暗黙的な定義を並べる書き方（複数の変数を定義する場合はvarで囲ったほうが望ましい）
	n := 1
	s := "string"
	b := true


}

func one() int {
	return 1
}
