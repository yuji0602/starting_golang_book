package foo

/*
1文字目が大文字の場合は他のパッケージから参照可能。
小文字の場合は他のパッケージから参照出来ない。
日本語などの大文字小文字がない文字である場合は他のパッケージから参照出来ない。
*/

/* 定数 */
const (
	MAX            = 100
	internal_const = 1
)

/* パッケージ変数 */
var (
	m = 256
	N = 512
)

/* 公開される関数 */
func FooFunc(n int) int {
	return internalFunc(n)
}

/* パッケージのみで参照できる関数 */
func internalFunc(n int) int {
	return n + 1
}
