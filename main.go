package main
/*
Problem statement: https://www.geeksforgeeks.org/sort-array-contain-1-n-values/
*/
import (
	"fmt"
)



func printArr(arr []int){
	for i:=0;i<len(arr);i++ {
		fmt.Printf("%d ", arr[i]);
	}
}
func main() {
	arr := []int{10, 7, 9, 2, 8,
                            3, 5, 4, 6, 1}
	printArr(arr);
	for i:=0; i< len(arr);i++ {
		arr[i] = i+1
	}
  fmt.Println("\nafter sorting")
	printArr(arr);
	// fmt.Println("Hello, playground")
	

}
