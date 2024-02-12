package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	var sell_count int
	var percent float64
	var price float64
	var percent_sum float64
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {

		price = 0
		percent = 0
		percent_sum = 0

		fmt.Fscan(in, &sell_count, &percent)

		for x := 0; x < sell_count; x++ {
			fmt.Fscan(in, &price)

			s := (price * percent) / 100

			r := float64(int(s))
			s = s - r
			percent_sum += s

		}

		percent_sum := fmt.Sprintf("%.2f", percent_sum)

		fmt.Println(percent_sum)
	}

}
