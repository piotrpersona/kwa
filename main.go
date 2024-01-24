package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseHelp(arg string) bool {
	arg = strings.ToLower(arg)
	switch arg {
	case "help", "h", "-h", "--help":
		return true
	}
	return false
}

func help() {
	fmt.Println(`kwa - super simple awk
Usage:
	
	cat stdin | kwa <column nr | default: 0>

kwa:
	o help - display this message.
`)
}

func main() {
	column := 0
	var err error
	if len(os.Args) == 2 {
		if parseHelp(os.Args[1]) {
			help()
			os.Exit(0)
		}
		column, err = strconv.Atoi(os.Args[1])
		if err != nil || column < 0 {
			fmt.Println("please provide a valid column number [0, ...]")
			os.Exit(1)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		splitted := strings.Fields(text)
		if len(splitted) <= column {
			return
		}
		fmt.Println(splitted[column])
	}
}
