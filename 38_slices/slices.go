package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	//最初の要素は含むが、最後の要素は除いた半開区間を選択
	var s []int = primes[1:4]
	fmt.Println(s)
}
