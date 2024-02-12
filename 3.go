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
	var str string
	var res bool
	var current_state []string
	current_state_p := &current_state
	last_cm := ""
	last_cm_p := &last_cm

	M := []string{"C", "R", "D"}
	R := []string{"C"}
	C := []string{"M"}
	D := []string{"M"}

	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {

		fmt.Fscan(in, &str)
		res = false
		*current_state_p = make([]string, 0)
		*last_cm_p = ""

		for k, v := range str {

			s := string(v)

			if k == 0 && s != "M" {
				break
			}

			if last_cm == s {
				res = false
				break
			}

			if len(*current_state_p) == 0 {

				if s == "M" {
					*current_state_p = M
				}
				if s == "R" {
					*current_state_p = R
				}
				if s == "C" {
					*current_state_p = C
				}
				if s == "D" {
					*current_state_p = D
				}
			}

			if *last_cm_p == "M" {
				res = check_commad(M, s)
			} else if *last_cm_p == "R" {
				res = check_commad(R, s)
			} else if *last_cm_p == "C" {
				res = check_commad(C, s)
			} else if *last_cm_p == "D" {
				res = check_commad(D, s)
			}

			*last_cm_p = s

			if k != 0 && !res {
				break
			}
		}

		if *last_cm_p != "D" {
			res = false
		}

		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

func check_commad(list []string, c string) bool {
	res := false

	for _, v := range list {
		if v == c {
			res = true
		}
	}

	return res
}
