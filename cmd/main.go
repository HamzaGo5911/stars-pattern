package main

import "fmt"

func main() {

	RightAngledTriangle()
	InverseRightAngledTriangle()
	Pyramid()
	InvertedPyramid()

}

func RightAngledTriangle() {
	rows := 5
	fmt.Println("right-angled triangle\n")

	for i := 0; i < rows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" ")
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func InverseRightAngledTriangle() {
	rows := 5
	fmt.Println("inverse of right-angled triangle\n")

	for i := 0; i < rows; i++ {
		for j := rows - 1; j >= i; j-- {
			fmt.Print(" ")
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func Pyramid() {
	rows := 5
	fmt.Println("Pyramid")

	for i := 0; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")

		}
		fmt.Println("")
	}
}

func InvertedPyramid() {
	fmt.Println("Inverted Pyramid\n")
	rows := 5

	for i := rows; i >= 1; i-- {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
