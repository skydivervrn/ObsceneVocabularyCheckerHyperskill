package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	tabooWordsFilename string
	sentenceToCheck    string
	tabooFileData      []byte
)

func main() {
	_, err := fmt.Scan(&tabooWordsFilename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = readFileContent(tabooWordsFilename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for {
		_, err = fmt.Scan(&sentenceToCheck)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		switch sentenceToCheck {
		case "exit":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			sentenceToPrint := ""
			for _, word := range strings.Split(sentenceToCheck, " ") {
				if checkWord(word) {
					stars := ""
					for range sentenceToCheck {
						stars += "*"
					}
					sentenceToPrint += stars + " "
				} else {
					sentenceToPrint += word
				}
			}
			fmt.Println(sentenceToPrint)
		}
	}
	//printFileContent()
}

func readFileContent(filename string) error {
	var err error
	tabooFileData, err = os.ReadFile(filename)
	if err != nil {
		return err
	}
	return nil
}

func checkWord(word string) bool {
	//fmt.Println(string(tabooFileData))
	lines := strings.Split(string(tabooFileData), "\n")
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), strings.ToLower(word)) {
			if word == "awesome" {
				fmt.Println(string(tabooFileData))
			}
			return true
		}
	}
	return false
}

func printFileContent() {
	fmt.Println(string(tabooFileData))
}
