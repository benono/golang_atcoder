package main

type Matrix [][]int

func mul(a, b *Matrix) Matrix {
	ar := len(*a)
	ac := len((*a)[0])
	br := len(*b)
	bc := len((*b)[0])

	// 縦横のサイズが合わない場合
	if ac != br {
		panic("wrong matrix type")
	}

	// ここが O(n^3) になってる
	c := make(Matrix, ar)
	for i := 0; i < ar; i++ {
		c[i] = make([]int, bc)
		for j := 0; j < bc; j++ {
			for k := 0; k < ac; k++ {
				c[i][j] += (*a)[i][k] * (*b)[k][j]
			}
		}
	}
	return c
}
