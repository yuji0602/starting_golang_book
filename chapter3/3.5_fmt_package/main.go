package main

import (
	"fmt"
)

func main() {
	/* fmt.Println */
	fmt.Println("Hello, Golang!")
	fmt.Println("My", "name", "is", "Taro")

	/* fmt.Printf */
	fmt.Printf("数値=%d\n", 5)
	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17)
	// パラメータが足りない（MISSINGが表示される）
	fmt.Printf("%d年%d月%d日\n", 2015, 12)
	// パラメータが多い（EXTRAが表示される）
	fmt.Printf("%d年%d月%d日\n", 2015, 12, 25, 17)
	fmt.Println("") // 上記で改行がおかしくなるので改行を追加
	// %vはさまざまな方のデータを埋め込む
	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3})
	// %#vはGoのリテラル表現でデータを埋め込む
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3})
	// %Tはデータの型情報を埋め込む
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3})

	/* print, println */
	/* printやprintlnは標準エラー出力へ出力される */
	print("Hello, World!")
	println("Hello, World!")
	println(1, 2, 3)
}
