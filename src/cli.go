package main

import (
	"bufio"
	"fmt"
	"gopass/commands"
	"os"
	"strings"
)

func main() {
	fmt.Println("hi, i'm a password manager cli written on golang.\nyou can easily store any password just with one command.\nyou can just write gopass and access to its cli, or use commands, like gopass help.")
	registeredCommands := commands.GenerateCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			input := strings.Split(scanner.Text(), " ")
			command := input[0]
			arguments := input[1:]
			output := commands.FindCorrespondingHandler(registeredCommands, command)(arguments)
			if output == "exit" {
				break
			}
			fmt.Println(output)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("bye")
	}
}
