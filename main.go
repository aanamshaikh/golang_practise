package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 1}
	fmt.Print(isDuplicate(number))

	fmt.Println(isAnagram("car", "rac"))

	fmt.Println(twoSum(number,3))
}
