package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkWord(tabooWords []string, inputWord string) bool {
	for _, tabooWord := range tabooWords {
		if tabooWord == strings.ToLower(inputWord) {
			//fmt.Println("True")
			return true
		}
	}
	return false
}

func censorWord(inputWord string) string {
	return strings.Repeat("*", len(inputWord))
}

func main() {
	// write your code here
	fileName := ""
	fmt.Scan(&fileName)
	tabooWords := []string{}
	inputWord := ""

	pInputScanner := bufio.NewScanner(os.Stdin)
	pInputScanner.Split(bufio.ScanWords)

	//fmt.Println(string(inputWord))

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pScanner := bufio.NewScanner(file)
	pScanner.Split(bufio.ScanWords)

	for pScanner.Scan() {
		tabooWords = append(tabooWords, pScanner.Text())
	}

	for pInputScanner.Scan() {
		inputWord = pInputScanner.Text()
		if inputWord == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		if checkWord(tabooWords, inputWord) {
			fmt.Print(censorWord(inputWord), " ")
		} else {
			fmt.Print(inputWord, " ")
		}
	}
	fmt.Println()

}
