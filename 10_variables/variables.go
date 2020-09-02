package main

import "fmt"

//var ステートメントは変数( variable )を宣言します。 関数の引数リストと同様に、複数の変数の最後に型を書くことで、変数のリストを宣言
var c, python, java bool

//var ステートメントはパッケージ、または、関数で利用できます。
func main() {
	var i int
	fmt.Println(i, c, python, java)
}
