package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	var n, m, ans int
	sc.Split(bufio.ScanWords)

	fmt.Scan(&n, &m)

	a := make([]int, n, n)

	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	for i := 0; i < m; i++ {
		ans += a[nextInt()-1]
	}

	fmt.Println(ans)
}
