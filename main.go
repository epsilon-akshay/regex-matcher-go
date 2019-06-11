package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("enter a word")
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')
	fmt.Println("enter the regex pattern you wish to input")
	regex, _ := reader.ReadString('\n')
	fmt.Printf("the word and the regex is %v %v", word, regex)
	var stack string
	compareString := getSetOfString(regex, stack)
	for i := 0; i < len(compareString); i++ {
		if regex[i] != compareString[i] {
			fmt.Println("false")
			return
		}
	}
	fmt.Println("true")
}

func getSetOfString(regex string, stack string) string {
	fmt.Println(len(regex), string(regex[3]))
	for i := 1; i < len(regex)-1; i = i + 1 {
		if regex[i-1] == '[' {
			for j := i; j < len(regex); j++ {
				stack = stack + string(regex[j])
				if regex[j] == ']' {
					i = j
					break
				}
			}
		}
	}
	return stack
}
