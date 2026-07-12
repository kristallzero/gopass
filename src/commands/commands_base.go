package commands

import "slices"

type command struct {
	aliases        []string
	description    string
	commandHandler func(arguments []string) string
}

func GenerateCommands() []command {
	commands := make([]command, 3)
	commands[0] = command{[]string{"help", "wtf", "what", "how"}, "write this message", HelpCommandHandler(commands)}
	commands[1] = command{[]string{"list", "all"}, "list of all stored passwords", ListCommandHandler}
	commands[2] = command{[]string{"exit", "quit", "bye", "gg", "e", "q"}, "exit", ExitCommandHandler}
	return commands
}

func FindCorrespondingHandler(commands []command, searchingCommand string) func(arguments []string) string {
	for _, command := range commands {
		if slices.Contains(command.aliases, searchingCommand) {
			return command.commandHandler
		}
	}
	return func(arguments []string) string { return "command not found" }
}
