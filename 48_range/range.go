package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//1つ目の変数はインデックス( index )で、2つ目はインデックスの場所の要素のコピー
	for i, v := range pow {
		fmt.Printf("2**%d = %d \n", i, v)
	}
}
