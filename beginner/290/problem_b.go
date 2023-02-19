package main

import "fmt"

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	var c int
	var s, ans string

	fmt.Scan(&s)

	for i, pass := range s {
		ans += string(pass)

		if string(pass) == string("o") {
			c++

			if c == k {
				for j := 0; j < (n - i - 1); j++ {
					ans += string("x")
				}
				break
			}
		}
	}

	fmt.Println(ans)
}
