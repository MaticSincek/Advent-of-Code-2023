package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	var cal_val_sum int = 0

	replacements := make(map[string]string)
	replacements["one"] = "1"
	replacements["two"] = "2"
	replacements["three"] = "3"
	replacements["four"] = "4"
	replacements["five"] = "5"
	replacements["six"] = "6"
	replacements["seven"] = "7"
	replacements["eight"] = "8"
	replacements["nine"] = "9"
	replacements["zero"] = "0"

	dat, err := os.Open("day1.txt")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		var line string = string(scanner.Text())
		var newline string = ""

		var numbertype int = 0
		var changed bool = true
		for changed {

			var mini int = 999
			var minmatching string = ""

			for key, _ := range replacements {
				i := strings.Index(line, key)
				if i < mini && i > -1 {
					mini = i
					minmatching = key
					numbertype = 1
				}
			}

			for _, value := range replacements {
				i := strings.Index(line, value)
				if i < mini && i > -1 {
					mini = i
					minmatching = value
					numbertype = 2
				}
			}

			if mini == 999 {
				changed = false
			} else {
				if numbertype == 1 {
					minmatching = replacements[minmatching]
				}
				line = line[mini+1:]
				newline = newline + minmatching
			}
		}

		var found_first, found_last bool = false, false
		var first_dig, last_dig string = "", ""
		for _, chr := range newline {

			if unicode.IsDigit(chr) {
				if !found_first {
					found_first = true
					first_dig = string(chr)
				} else {
					found_last = true
					last_dig = string(chr)
				}
			}
		}

		var cal_string string = "0"

		if found_first && !found_last {
			cal_string = first_dig + first_dig
		} else if found_first && found_last {
			cal_string = first_dig + last_dig
		} else {
			cal_string = "0"
		}

		cal_val, err := strconv.Atoi(cal_string)
		if err != nil {
			fmt.Printf("Error converting to int: %s\n", err)
			os.Exit(-2)
		}

		cal_val_sum += cal_val
	}

	fmt.Printf("%s\n", strconv.Itoa(cal_val_sum))

	dat.Close()
}
