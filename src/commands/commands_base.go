package commands

import (
	"fmt"
	"gopass/data"
	"slices"
)

type command struct {
	aliases        []string
	description    string
	commandHandler func(*data.Storage, []string) string
}

func GenerateCommands() []command {
	commands := make([]command, 4)
	commands[0] = command{[]string{"help", "wtf", "what", "how"}, "write this message", HelpCommandHandler(commands)}
	commands[1] = command{[]string{"list", "all", "read", "get", "access"}, "list of all stored credentials. write a source name as an argument to get only the source's credentials. if there are more than one credential per source, you can also provide the login", ListCommandHandler}
	commands[2] = command{[]string{"add", "create", "bye", "new"}, "add (<source> <login> <password>) add new credentials. source is a place where the credentials are being used", AddCommandHandler}
	commands[3] = command{[]string{"exit", "quit", "bye", "gg", "e", "q"}, "exit", ExitCommandHandler}
	return commands
}

func FindCorrespondingHandler(commands []command, searchingCommand string) func(*data.Storage, []string) string {
	for _, command := range commands {
		if slices.Contains(command.aliases, searchingCommand) {
			return command.commandHandler
		}
	}
	return func(*data.Storage, []string) string { return "command not found" }
}

func checkArgumentsLength(arguments []string, expectedLength int) string {
	actualLength := len(arguments)
	if expectedLength == actualLength {
		return ""
	}
	return fmt.Sprintf("incorrect arguments length. expected: %d, got: %d", expectedLength, actualLength)
}

func checkArgumentsLengthMaximum(arguments []string, expectedMaxLength int) string {
	actualLength := len(arguments)
	if actualLength <= expectedMaxLength {
		return ""
	}
	return getArgumentsLengthMaximumMessage(actualLength, expectedMaxLength)
}

func getArgumentsLengthMaximumMessage(actualLength, expectedMaxLength int) string {
	return fmt.Sprintf("incorrect arguments length. expected: 0-%d, got: %d", expectedMaxLength, actualLength)
}
