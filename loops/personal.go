package main

import "fmt"

func main() {

	var rockslice []string
	rockslice = append(rockslice, "Cyndi", "Lauper", "69")
	rockslice = append(rockslice, "Elvis", "Presley", "87")
	rockslice = append(rockslice, "Bruce", "Springsteen", "73")
	rockslice = append(rockslice, "Billy", "Billy", "67")
	rockslice = append(rockslice, "Ozzy", "Osbourne", "74")
	rockslice = append(rockslice, "Brandon", "Flowers", "41")

	fmt.Println("List available data to upload")
	fmt.Println()

	// Simply print a slice in a loop with a breakline very 3 elements
	fmt.Println("------------METHOD I------------")
	n := 1
	for i := 0; i <= len(rockslice)-1; i++ {
		fmt.Print(rockslice[i], " ")
		if n == 3 {
			fmt.Println()
			n = 0
		}
		n++
	}

	// Build a new slice with 3 elements then print
	fmt.Println("------------METHOD II-----------")
	var readyslice []string
	r := 1
	for i := 0; i <= len(rockslice)-1; i++ {
		readyslice = append(readyslice, rockslice[i])
		if r == 3 {
			fmt.Println("--> ", readyslice[0], readyslice[1], readyslice[2])
			r = 0
			readyslice = nil
		}
		r++
	}

}
