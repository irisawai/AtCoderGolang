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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	y := nextInt()

	var b1000, b5000, b10000 int

	for b10000 = 0; b10000 <= n; b10000++ {
		if b10000*10000 > y {
			break
		}

		for b5000 = 0; b5000 <= (n - b10000); b5000++ {
			if b10000*10000+b5000*5000 > y {
				break
			}

			b1000 = n - b10000 - b5000

			if b10000*10000+b5000*5000+b1000*1000 == y {
				goto L1
			}
		}
	}

	b1000 = -1
	b5000 = -1
	b10000 = -1

L1:
	fmt.Println(b10000, b5000, b1000)
}
