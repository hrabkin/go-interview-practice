package main

import (
	"fmt"
)

func main() {
	// Example slice for testing
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// Test FindMax
	max := FindMax(numbers)
	fmt.Printf("Maximum value: %d\n", max)

	// Test RemoveDuplicates
	unique := RemoveDuplicates(numbers)
	fmt.Printf("After removing duplicates: %v\n", unique)

	// Test ReverseSlice
	reversed := ReverseSlice(numbers)
	fmt.Printf("Reversed: %v\n", reversed)

	// Test FilterEven
	evenOnly := FilterEven(numbers)
	fmt.Printf("Even numbers only: %v\n", evenOnly)
}

// FindMax returns the maximum value in a slice of integers.
// If the slice is empty, it returns 0.
func FindMax(numbers []int) int {
    
	s := len(numbers)
	if s == 0 { return 0 }
	
	msf := numbers[0]
	for _,v := range(numbers) {
	    if v > msf {
	        msf = v
	    }
	}
	
	return msf
}

// RemoveDuplicates returns a new slice with duplicate values removed,
// preserving the original order of elements.
func RemoveDuplicates(numbers []int) []int {
    
	seen := make(map[int]bool)
	res := make([]int, 0, len(numbers))
	
	for _,v := range(numbers) {
	    
	    if _, exists := seen[v]; !exists {
	        seen[v] = true
	        res = append(res, v)   
	    }
	}
	
	return res
}

// ReverseSlice returns a new slice with elements in reverse order.
func ReverseSlice(slice []int) []int {
    
    l,r := 0, len(slice)-1
	res := make([]int, len(slice))
	copy(res, slice)
	
    for l < r {
        res[l] = slice[r]
        res[r] = slice[l]
        l += 1
        r -= 1
    }
    
	return res
}

// FilterEven returns a new slice containing only the even numbers
// from the original slice.
func FilterEven(numbers []int) []int {
    
	res := make([]int, 0, len(numbers))
	
	for _,v := range(numbers) {
	    if v % 2 == 0 {
	        res = append(res, v)
	    }
	}
	
	return res
}
