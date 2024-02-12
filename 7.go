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

	var olde_count int
	var olde string
	olde_list := make(map[string]bool, 0)
	var newe_count int
	var newe string
	//newe_list := make([]string, 0)

	//var str string
	fmt.Fscan(in, &olde_count)

	for i := 0; i < olde_count; i++ {
		fmt.Fscan(in, &olde)
		olde_list[olde] = true
	}

	fmt.Fscan(in, &newe_count)

	for i := 0; i < newe_count; i++ {
		fmt.Fscan(in, &newe)
		res := check_login(olde_list, newe)
		fmt.Println(res)
	}

	//fmt.Println(olde_list)
	//fmt.Println(newe_list)

}

func check_login(olde_list map[string]bool, newe string) int {
	//check := 0

	if olde_list[newe] {
		return 1
	}

	newb := []byte(newe)

	for i := 0; i < len(newb)-1; i++ {
		newb[i], newb[i+1] = newb[i+1], newb[i]
		if olde_list[string(newb)] {
			return 1
		}

		newb[i], newb[i+1] = newb[i+1], newb[i]
	}
	return 0
}
