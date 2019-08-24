package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	for { // Runs forever until ctrl+c or we 'return' from the loop
		fmt.Printf(
			"Hello! What do you want to do?:\n" + "" +
				"1. Register.\n" +
				"2. Sign in.\n" +
				"3. Quit.\n")
		// Reads a single UTF-8 character (rune)
		// from STDIN and switches to case.
		switch ask() {
		case "1":
			fmt.Println("Registering new account!")
			break // 'breaks' out of the switch, but we're still in the 'for' loop.
		case "2":
			fmt.Println("Signing in!")
			break
		case "3":
			fmt.Println("Bye!")
			return // quits the loop
		default:
			fmt.Println("Invalid option. Please try again.")
			break
		}
	}
}

func ask() string {
	reader := bufio.NewReader(os.Stdin)
	inputVal, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("invalid option: %v\n", err)
		return ""
	}

	output := strings.TrimSuffix(inputVal, "\n") // Important!
	return output
}

func loadCSV(filename string) (interface{}, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
}
