package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type KtaFolder struct {
	Dir     string       `json:"dir"`
	Files   []string     `json:"files,omitempty"`
	Folders []*KtaFolder `json:"folders,omitempty"`
}

var hack_sum int

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var lines int

	fmt.Fscan(in, &lines)

	for i := 0; i < lines; i++ {

		var lines_count int
		var fullstr_json strings.Builder
		fmt.Fscan(in, &lines_count)

		for j := -1; j < lines_count; j++ {
			line, err := in.ReadString('\n')
			if err != nil {
			}
			fullstr_json.WriteString(line)
		}

		var result KtaFolder

		json.Unmarshal([]byte(fullstr_json.String()), &result)

		hack_sum = run(&result, false)

		fmt.Println(hack_sum)
	}

}

func files_count(files []string) (int, bool) {

	var has_hacked_files bool
	for _, file := range files {
		if len(file) > 5 && file[len(file)-5:] == ".hack" {
			has_hacked_files = true
		}
	}

	return len(files), has_hacked_files
}

func run(folder *KtaFolder, parent_has_hacked_files bool) int {
	res := 0

	files_count, has_hacked_files := files_count(folder.Files)
	has_hacked_files = has_hacked_files || parent_has_hacked_files
	if has_hacked_files {
		res += files_count
	}

	if len(folder.Folders) != 0 {
		for _, subfolder := range folder.Folders {
			res += run(subfolder, has_hacked_files)
		}
	}
	return res
}
