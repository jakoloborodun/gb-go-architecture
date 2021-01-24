package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"

	util "go_arch/lesson-4/pkg/random"
)

func main() {
	searchArray := util.RandomIntSlice(10)
	sort.Ints(searchArray)

	var searchKey int
	fmt.Printf("Search array is %v. What is your search key?\n\n", searchArray)
	fmt.Scanf("%d\n", &searchKey)

	result, err := binarySearch(searchKey, searchArray)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//result := binaryRecursive(searchKey, searchArray)

	if result < 0 {
		fmt.Printf("Your key %d was not found in the array %v.\n\n", searchKey, searchArray)
	} else {
		fmt.Printf("Your key was found in position %d in the array %v.\n\n", result, searchArray)
	}
}

func binarySearch(key int, array []int) (position int, err error) {
	position = -1

	if len(array) == 0 {
		return position, errors.New("empty search array")
	}

	first := 0
	last := len(array)
	for first < last {
		mid := first + (last-first)/2
		if array[mid] == key {
			position = mid
			break
		} else {
			if array[mid] > key {
				last = mid
			} else {
				first = mid + 1
			}
		}
	}

	return position, nil
}

//
//func binaryRecursive(key int, array []int) (position int) {
//	mid := len(array) / 2 // mid should be int cause we use 2 int digits
//	switch {
//	case len(array) == 0:
//		position = -1 // not found
//	case array[mid] > key:
//		position = binaryRecursive(key, array[:mid])
//	case array[mid] < key:
//		position = binaryRecursive(key, array[mid+1:])
//		if position >= 0 { // if anything but the -1 "not found" result
//			position += mid + 1
//		}
//	default:
//		position = mid // found
//	}
//
//	return position
//}
