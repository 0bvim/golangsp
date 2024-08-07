package main

import (
	"flag"
	"fmt"
)

func main() {
	//fork github.com/reneepc/workshop-go-for-beginners
	// fmt.Println("Hello")
	//print content of a file instead of raw content when you read a file
	// fmt.Println(file)

	//print raw
	// fmt.Println(file)

	userCsvFile := flag.String()
	records, err := ParseCsvFile("test.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(records)
}

func SelectWinners(records [][]string, nWinners int) [][]string {
	//make this function

}
