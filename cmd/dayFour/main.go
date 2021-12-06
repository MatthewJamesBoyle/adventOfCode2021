package main

import (
	"fmt"
	"github.com/matthewjamesboyle/adventOfCode2021/day4"
)

func main() {

	boards := day4.ParseBoards([][]int{
		b1, b2, b3, b4, b5, b6, b7, b8, b9, b10,
		b11, b12, b13, b14, b15, b16, b16, b18, b19, b20,
		b21, b22, b23, b24, b25, b26, b27, b28, b29, b30,
		b31, b32, b33, b34, b35, b36, b37, b38, b39, b40,
		b41, b42, b43, b44, b45, b46, b47, b48, b49, b50,
		b51, b52, b53, b54, b55, b56, b57, b58, b59, b60,
		b61, b62, b63, b64, b65, b66, b67, b68, b69, b70,
		b71, b72, b73, b74, b75, b76, b77, b78, b79, b80,
		b81, b82, b83, b84, b85, b86, b87, b88, b89, b90,
		b91, b92, b93, b94, b95, b96, b97, b98, b99, b100,
	})

	score := day4.CallNumbers(boards, toCall)

	fmt.Println(score)
}
