package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FontAscii(value, font string) (string, int) {

	file, eR := OpenFileAndCheckFont(font)
	if eR == 500 {
		return "", 500
	}
	// func scanner, openned
	fileBaseLetter := ScanFile(file)
	ans := ""
	narg := strings.Split(value, "\\n")

	if font == "shadow" || font == "standard" || font == "thinkertoy" {
		for _, v := range narg {

			for i := 1; i <= 8; i++ {
				for _, arg := range v {
					//find index, args letter, where locate ascii letter - fileVal array
					indexBase := int(rune(arg)-32) * 9
					ans += fmt.Sprint(fileBaseLetter[indexBase+i])
				}
				// fmt.Println()
				ans += "\n"
			}
		}

	} else {
		fmt.Println("Такого шрифта не найдена")
	}
	return ans, 200
}
func printLetter(s string, fileBaseLetter []string) string {
	ans := ""
	for i := 1; i <= 8; i++ {
		for _, arg := range s {
			//find index, args letter, where locate ascii letter - fileVal array
			indexBase := int(rune(arg)-32) * 9
			ans += fmt.Sprint(fileBaseLetter[indexBase+i])
		}
		// fmt.Println()
		ans += "\n"
	}

	return ans
}

//scan open file, return data
func ScanFile(arg *os.File) []string {
	var fileValue []string
	scanner := bufio.NewScanner(arg)
	for scanner.Scan() {
		lines := scanner.Text()
		fileValue = append(fileValue, lines)
	}
	return fileValue
}

//open font file, and  0, 0, 0return hims
func OpenFileAndCheckFont(input string) (*os.File, int) {
	file, err := os.Open("source/"+input+".txt")
	if err != nil {
		return nil, 500
	}
	return file, 200
}
