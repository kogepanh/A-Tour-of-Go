package main

import "fmt"

func main() {
	//defer ステートメントは、 defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させる
	defer fmt.Println("world")

	fmt.Println("hello")
}
