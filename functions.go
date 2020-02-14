package main

import (
	"bufio"
	"os"
)

func handleNewLine(str string) []string {
	var strs []string
	last := 0
	for i := 0; i < len(str); i++ {
		if i == len(str)-1 {
			strs = append(strs, str[last:i+1])
		}
		if str[i] == '\\' && str[i+1] == 'n' {
			strs = append(strs, str[last:i])
			last = i + 2
		}

	}
	return strs
}

func readLine(lineNum int, filename string, result *string) {
	file, _ := os.Open(filename)
	defer file.Close()
	sc := bufio.NewScanner(file)
	lastLine := 0

	for sc.Scan() {
		if lastLine == lineNum {
			*result = *result + sc.Text()
			// fmt.Printf(sc.Text())
			break
		}
		lastLine++
	}
	*result = *result + ""
	// fmt.Printf("")
}

func drawBanner(args []string, result *string) {

	strs := handleNewLine(args[0])
	var lineNum int
	var nxt int
	var filename string

	if len(args) > 1 {
		filename = args[1] + ".txt"
	} else {
		filename = "standard.txt"
	}

	for n := 0; n < len(strs); n++ {
		nxt = 0
		if len(strs[n]) > 0 {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(strs[n]); j++ {
					lineNum = (int(strs[n][j])-32)*9 + 1
					readLine(lineNum+nxt, filename, result)
				}
				nxt++
				// fmt.Println()
				*result = *result + "\n"
			}
		} else {
			// fmt.Println()
			*result = *result + "\n"
		}
	}
}

func Font(args []string) string {
	if len(args) > 0 {
		result := ""
		drawBanner(args, &result)
		return result
	}
	return ""
}
