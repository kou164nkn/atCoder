package main

import (
	"fmt"
)

func main() {
	var n, m, ans int
	for {
		fmt.Scan(&n, &m)

		if (1 <= n && n <= 100) && (1 <= m && m <= 100) && (m < n) {
			break
		}
	}

	a := make([]int, n, n)
	b := make([]int, m, m)

	for i := 0; i < n; i++ {
		var tmp int

		for {
			fmt.Scan(&tmp)

			if (1 <= tmp) && (tmp <= 100) {
				break
			}
		}

		a[i] = tmp
	}

	for i := 0; i < m; i++ {
		var tmp int

		for {
			fmt.Scan(&tmp)

			if (1 <= tmp) && (tmp <= 100) && (tmp < a[i]) {
				break
			}
		}
		b[i] = tmp
	}

	for _, point := range b {
		ans += point
	}

	fmt.Println(ans)
}
