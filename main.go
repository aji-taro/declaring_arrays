package main

import (
	"fmt"
)

func main() {
	fmt.Println("# go")

	// 基本
	array01 := [3]int{1, 2, 3}
	fmt.Println("array01", array01, "length:", len(array01))

	// ... を指定すると要素数を自動で設定してくれる
	array02 := [...]int{10, 20, 30, 40}
	fmt.Println("array02", array02, "length:", len(array02))

	// 要素数を省略するとスライスになる
	array03 := []int{100, 200, 300, 400, 500}
	fmt.Println("array03", array03, "length:", len(array03))

	// make でスライスを作成
	array04 := make([]int, 6)
	// for i := range array04 {
	for i := 0; i < len(array04); i++ {
		array04[i] = i + 1
	}
	fmt.Println("array04", array04, "length:", len(array04))

	// 二次元配列 (2行 * 3列の配列の例)
	array05 := [2][3]int{
		{1, 2, 3},
		{10, 20, 30},
	}
	// あえてループしてみる
	for i := range array05 {
		fmt.Print("array05 [index:", i, "]")
		for _, v := range array05[i] {
			fmt.Print(" ", v)
		}
		fmt.Print("\n")
	}

	// 要素の追加
	a := []int{11, 22, 32}
	array06 := append(a, 44, 55)
	fmt.Println("array06", array06, "length:", len(array06))

	// 連想配列 (map)
	array07 := map[string]int{
		"goals":   3,
		"assists": 5,
		"duels":   7,
	}
	fmt.Println("array07", array07, "length:", len(array07))
}
