package main

import "fmt"

func main() {
	var n, awal, naik int
	fmt.Scan(&n, &awal, &naik)
	total := awal
	naik_ke := naik
	for i := 1; i <= n; i++ {
		total += naik_ke
		naik_ke += naik
	}
	fmt.Println(total)
}
