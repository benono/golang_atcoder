package main

import "fmt"

func main() {

}

// 1_6_D
// 入力: n行m列のベクトル, n
func matrixVectorMultiplication() {
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	b := make([]int, m)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &b[i])
	}

	for i := 0; i < n; i++ {
		var sum int
		for j := 0; j < m; j++ {
			sum += a[i][j] * b[j]
		}
		fmt.Println(sum)
	}
}

// 1_7_A 成績評価
func grading() {
Loop:
	for {
		var m int
		var f int
		var r int
		fmt.Scan(&m)
		fmt.Scan(&f)
		fmt.Scan(&r)

		var grade string
		if (m == -1) && (f == -1) && (r == -1) {
			break Loop

		}
		// 成績評価：中間試験と期末試験の合計点
		sum := m + f
		if sum >= 80 {
			grade = "A"
		} else if sum >= 65 {
			grade = "B"
		} else if sum >= 50 {
			grade = "C"
		} else if sum >= 30 {
			if r >= 50 {
				grade = "C"
			} else {
				grade = "D"
			}
		} else {
			grade = "F"
		}

		if (m == -1) || (f == -1) {
			grade = "F"

		}
		fmt.Println(grade)

	}
}

// 1_7_B
func howManyWays() {
	for {
		var n, x int
		fmt.Scanf("%d %d", &n, &x)
		if n == 0 && x == 0 {
			break
		}
		var cnt int
		for i := 1; i < n; i++ {
			for j := i + 1; j <= n; j++ {
				if j == i {
					continue
				}
				for k := j + 1; k <= n; k++ {
					if k == i || k == j {
						continue
					}
					if (i + j + k) == x {
						cnt++
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}

// 1_7_C
func spreadSheet() {
	var r, c int
	fmt.Scanf("%d %d", &r, &c)
	sheet := make([][]int, r)
	for i := 0; i < r; i++ {
		sheet[i] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scanf("%d", &sheet[i][j])
		}
	}
	for i := 0; i < r; i++ {
		var rowCnt int
		for j := 0; j < c; j++ {
			fmt.Printf("%d ", sheet[i][j])
			rowCnt += sheet[i][j]
		}
		fmt.Println(rowCnt)
	}

	var cnt int
	for i := 0; i < c; i++ {
		var colCnt int
		for j := 0; j < r; j++ {
			colCnt += sheet[j][i]
		}
		fmt.Printf("%d ", colCnt)
		cnt += colCnt
	}
	fmt.Println(cnt)
}
