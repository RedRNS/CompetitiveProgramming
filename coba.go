package main

import "fmt"

func main() {
	var n, total, pembagi int64
	total, pembagi = 0, 0
	fmt.Scan(&n)

	for {
		if n == 9999 {
			break
		}
		pembagi++
		total += n
		fmt.Scan(&n)
	}
	fmt.Print(total/pembagi + 2)
}
