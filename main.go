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
	fmt.Printf("%v %v", word, regex)
}
