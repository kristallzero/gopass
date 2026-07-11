package main

import "slices"

type Command struct {
	aliases        []string
	description    string
	commandHandler func(arguments []string) string
}

func GenerateCommands() []Command {
	commands := make([]Command, 3)
	commands[0] = Command{[]string{"help", "wtf", "what", "how"}, "write this message", HelpCommandHandler(commands)}
	commands[1] = Command{[]string{"list", "all"}, "list of all stored passwords", ListCommandHandler}
	commands[2] = Command{[]string{"exit", "quit", "bye", "gg", "e", "q"}, "exit", ExitCommandHandler}
	return commands
}

func FindCorrespondingHandler(commands []Command, searchingCommand string) func(arguments []string) string {
	for _, command := range commands {
		if slices.Contains(command.aliases, searchingCommand) {
			return command.commandHandler
		}
	}
	return func(arguments []string) string { return "command not found" }
}
