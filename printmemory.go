package main

import "fmt"

func PrintMemory(arr [10]byte) {
	for i := 0; i < len(arr); i += 4 {
		end := i + 4
		if end > len(arr) {
			end = len(arr)
		}
		for j := i; j < end; j++ {
			if j > i {
				fmt.Print(" ")
			}
			fmt.Printf("%02x", arr[j])
		}
		fmt.Println()
	}
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
	fmt.Println()
}
