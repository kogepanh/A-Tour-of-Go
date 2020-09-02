package main

import "fmt"

// naked returnステートメントは、短い関数でのみ利用すべき
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
