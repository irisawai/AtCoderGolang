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

func divideDegit(d int) [5]int {
	var d_ar [5]int

	d_ar[0] = d % 10
	d_ar[1] = (d % 100) / 10
	d_ar[2] = (d % 1000) / 100
	d_ar[3] = (d % 10000) / 1000
	d_ar[4] = d / 10000

	return d_ar
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := nextInt()
	b := nextInt()
	sum := 0

	d_ar := divideDegit(n)
	var d5max, d4max, d3max, d2max, d1max int

	d5max = Min(d_ar[4], b)
	if d5max > 9 {
		d5max = 9
	}

	for d5 := 0; d5 <= d5max; d5++ {
		if d5 == d_ar[4] {
			d4max = Min(d_ar[3], b-d5)
		} else {
			d4max = b - d5
		}

		if d4max > 9 {
			d4max = 9
		}

		for d4 := 0; d4 <= d4max; d4++ {
			if d5 == d_ar[4] && d4 == d_ar[3] {
				d3max = Min(d_ar[2], b-d5-d4)
			} else {
				d3max = b - d5 - d4
			}

			if d3max > 9 {
				d3max = 9
			}

			for d3 := 0; d3 <= d3max; d3++ {
				if d5 == d_ar[4] && d4 == d_ar[3] && d3 == d_ar[2] {
					d2max = Min(d_ar[1], b-d5-d4-d3)
				} else {
					d2max = b - d5 - d4 - d3
				}

				if d2max > 9 {
					d2max = 9
				}

				for d2 := 0; d2 <= d2max; d2++ {
					if d5 == d_ar[4] && d4 == d_ar[3] && d3 == d_ar[2] && d2 == d_ar[1] {
						d1max = Min(d_ar[0], b-d5-d4-d3-d2)
					} else {
						d1max = b - d5 - d4 - d3 - d2
					}

					if d1max > 9 {
						d1max = 9
					}

					for d1 := 0; d1 <= d1max; d1++ {
						if d1+d2+d3+d4+d5 >= a {
							sum += d1 + d2*10 + d3*100 + d4*1000 + d5*10000
						}
					}
				}
			}
		}
	}

	fmt.Println(sum)
}
