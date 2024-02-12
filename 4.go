package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cursor struct {
	pos_x, pos_y int
	name         string
	queue        int
}

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	var x, y int
	var str string

	a_cursor := &Cursor{
		pos_x: 0,
		pos_y: 0,
		name:  "a",
		queue: 0,
	}

	b_cursor := &Cursor{
		pos_x: 0,
		pos_y: 0,
		name:  "b",
		queue: 0,
	}

	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {

		fmt.Fscan(in, &y, &x)

		a_cursor.pos_x = 0
		a_cursor.pos_y = 0
		a_cursor.queue = 0
		b_cursor.pos_x = 0
		b_cursor.pos_y = 0
		b_cursor.queue = 0

		field_list := make([][]string, y)

		for c := 0; c < y; c++ {
			fmt.Fscan(in, &str)

			for _, v := range str {
				field_list[c] = append(field_list[c], string(v))
			}

		}

		for kyy, yy := range field_list {
			strs := ""
			for kxx, xx := range yy {
				if xx == "A" {
					a_cursor.pos_y = kyy
					a_cursor.pos_x = kxx
					if b_cursor.queue == 0 {
						a_cursor.queue = 1
					}

				}
				if xx == "B" {
					b_cursor.pos_y = kyy
					b_cursor.pos_x = kxx

					if a_cursor.queue == 0 {
						b_cursor.queue = 1
					}
				}
				strs += xx
			}
		}

		// fmt.Println("A.x ", a_cursor.pos_x)
		// fmt.Println("A.y ", a_cursor.pos_y)
		// fmt.Println("B.x ", b_cursor.pos_x)
		// fmt.Println("B.y ", b_cursor.pos_y)

		//	print_list(field_list)

		for field_list[0][0] == "." || field_list[y-1][x-1] == "." {

			if a_cursor.queue == 1 {
				field_list = move_up(field_list, a_cursor)
				field_list = move_left(field_list, a_cursor)
			} else {
				field_list = move_right(field_list, a_cursor)
				field_list = move_down(field_list, a_cursor)
			}

			if b_cursor.queue == 1 {
				field_list = move_up(field_list, b_cursor)
				field_list = move_left(field_list, b_cursor)
			} else {
				field_list = move_right(field_list, b_cursor)
				field_list = move_down(field_list, b_cursor)
			}

		}

		print_list(field_list)

		// fmt.Println("A.x ", a_cursor.pos_x)
		// fmt.Println("A.y ", a_cursor.pos_y)
		// fmt.Println("B.x ", b_cursor.pos_x)
		// fmt.Println("B.y ", b_cursor.pos_y)

		// for y, vv := range field_list {
		// 	strs := ""
		// 	for x, v := range vv {

		// 		strs += v
		// 	}
		// 	fmt.Println(strs)
		// }

	}

}

func print_list(field_list [][]string) {
	for _, yy := range field_list {
		strs := ""
		for _, xx := range yy {
			strs += xx
		}
		fmt.Println(strs)
	}
}

func move_up(field_list [][]string, cursor *Cursor) [][]string {

	if cursor.pos_y != 0 {
		if cursor.pos_y >= 0 {
			if (field_list[cursor.pos_y-1][cursor.pos_x]) == "." {
				field_list[cursor.pos_y-1][cursor.pos_x] = cursor.name
				cursor.pos_y--
			}
		}
	}

	return field_list
}

func move_left(field_list [][]string, cursor *Cursor) [][]string {

	if cursor.pos_x != 0 {
		if cursor.pos_x-1 >= 0 {
			if field_list[cursor.pos_y][cursor.pos_x-1] == "." {
				field_list[cursor.pos_y][cursor.pos_x-1] = cursor.name
				cursor.pos_x--
			}
		}
	}

	return field_list
}

func move_right(field_list [][]string, cursor *Cursor) [][]string {

	if cursor.pos_x != len(field_list[cursor.pos_y])-1 {

		if cursor.pos_x+1 <= len(field_list[cursor.pos_y])-1 {

			if field_list[cursor.pos_y][cursor.pos_x+1] == "." {
				field_list[cursor.pos_y][cursor.pos_x+1] = cursor.name
				cursor.pos_x++
			}

		}
	}

	return field_list
}

func move_down(field_list [][]string, cursor *Cursor) [][]string {

	if cursor.pos_y != len(field_list)-1 {

		if cursor.pos_y+1 <= len(field_list)-1 {

			if field_list[cursor.pos_y+1][cursor.pos_x] == "." {
				field_list[cursor.pos_y+1][cursor.pos_x] = cursor.name
				cursor.pos_y++
			}

		}

	}

	return field_list
}
