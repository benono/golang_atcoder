package main

import "fmt"

// 1_7_Cの模範解答
func spreadSheetMA() {
	var r, c int
	fmt.Scanf("%d %d", &r, &c)
	sheet := make([][]int, r+1)
	for i := 0; i <= r; i++ {
		sheet[i] = make([]int, c+1)
	}
	// slice作成時に、合計値もいれている。スゲー
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scanf("%d", &sheet[i][j])
			sheet[r][j] += sheet[i][j]
			sheet[i][c] += sheet[i][j]
			sheet[r][c] += sheet[i][j]
		}
	}

	for i := 0; i <= r; i++ {
		for j := 0; j <= c; j++ {
			fmt.Print(sheet[i][j])
			if j != c {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
