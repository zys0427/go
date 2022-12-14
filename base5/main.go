/**
go  切片
	删除指定KEY的切片元素
	添加切片的元素
*/
package main

import "fmt"

// 通过替换的方式 删除元素  会打乱数组元素
func remove(slice []int, i int) []int {
	n := len(slice)
	arr := make([]int, n)
	copy(arr, slice[:n])
	if i == n {
		return arr[0:n-1]
	}
	//fmt.Println(len(slice))
	arr[i] = arr[len(arr)-1]
	////fmt.Println(slice)
	return arr[:len(arr)-1]
	//return arr
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println("--1--", s, s[:4])
	s = []int{5, 6, 7, 8, 9}
	fmt.Println("--2--", s, remove(s, 2)) // "[5 6 9 8]
	s = []int{5, 6, 7, 8, 9}
	fmt.Println("--3--", s, addSlice(s, 99))
	s = []int{5, 6, 7, 8, 9}
	fmt.Println("--4--", s, removeSlice(s, 1))
	s = []int{5, 6, 7, 8, 9}
	fmt.Println("--5--", s, removeSliceTwo(s, 1))
}

func addSlice(sliceArr []int, i int) []int {
	//append(slice, i)
	sliceArr = append(sliceArr, i)
	return sliceArr
}

// 通过循环删除元素
func removeSlice(sliceArr []int, i int) []int {
	arr := make([]int, 0)
	if len(sliceArr) < 0 {
		return arr
	}
	for index, value := range sliceArr {
		if i != index {
			arr = append(arr, value)
		}
	}
	return arr
}

// 通过 从删除的元素 切开，然后分别添加到新切片 不会打乱元素顺序
func removeSliceTwo(sliceArr []int, i int) []int {
	arr := make([]int, i)
	copy(arr,sliceArr[0:i])
	arr = append(arr, sliceArr[i+1:]...)
	return arr
}