package main

import "fmt"

func main() {
	var n, m, ans int

	for {
		fmt.Scan(&n, &m)

		if (0 < n && n <= 1000) && (0 < m && m <= 1000) {
			break
		}
	}

	s := make([]string, n, n)
	t := make([]string, m, m)

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		s[i] = tmp[len(tmp)-3:]
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&t[i])
	}

	for _, num_s := range s {
		for _, num_t := range t {
			if num_s == num_t {
				ans++
				break
			}
		}
	}

	fmt.Println(ans)
}
