package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	column := 0
	var err error
	if len(os.Args) == 2 {
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
