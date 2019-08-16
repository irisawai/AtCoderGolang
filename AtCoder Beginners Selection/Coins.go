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

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()
	c := nextInt()
	x := nextInt()

	var remainAf500, remainAf100, max100 int
	cnt := 0
	max500 := Min(a, x/500)

	for i := 0; i <= max500; i++ {
		remainAf500 = x - 500*i
		max100 = Min(b, remainAf500/100)

		for j := 0; j <= max100; j++ {
			remainAf100 = remainAf500 - 100*j

			if remainAf100/50 <= c {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
