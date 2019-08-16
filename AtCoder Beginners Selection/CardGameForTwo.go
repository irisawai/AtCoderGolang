package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	a := make([]int, n, n)

	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	sort.Sort(sort.IntSlice(a))

	alice := 0
	bob := 0
	i := len(a) - 1
	for i >= 0 {
		alice += a[i]
		i--
		if i >= 0 {
			bob += a[i]
		} else {
			break
		}
		i--
	}

	fmt.Println(alice - bob)

}
