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

func devides(ar []int) (result []int, isDone bool) {
	for i := 0; i < len(ar); i++ {
		if ar[i]%2 == 0 {
			ar[i] = ar[i] / 2
		} else {
			return ar, false
		}
	}
	return ar, true
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	ar := make([]int, n, n)

	for i := 0; i < n; i++ {
		ar[i] = nextInt()
	}

	isDone := false
	cnt := 0

	for true {
		ar, isDone = devides(ar)

		if isDone {
			cnt += 1
		} else {
			break
		}
	}

	fmt.Println(cnt)
}
