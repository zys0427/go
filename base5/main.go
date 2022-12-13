package main

import "fmt"

func remove(slice []int, i int) []int {
	fmt.Println(len(slice))
	slice[i] = slice[len(slice)-1]
	fmt.Println(slice)
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(s[:4])
	fmt.Println(remove(s, 2)) // "[5 6 9 8]
}