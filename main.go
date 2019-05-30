package main

<<<<<<< HEAD
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
	fmt.Printf("the word and the regex is %v %v", word, regex)import "fmt"

func main() {
	fmt.Println("vim-go")
}
