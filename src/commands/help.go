package commands

import "strings"

func HelpCommandHandler(commands []Command) func(arguments []string) string {
	return func(arguments []string) string {
		var result strings.Builder
		result.WriteString("gopass <command>\nor\ngopass\n> <command>\n\ncommands:\n")

		for _, command := range commands {
			for i, alias := range command.aliases {
				result.WriteString(alias)
				if i != len(command.aliases)-1 {
					result.WriteString(" / ")
				}
			}
			result.WriteString(" - ")
			result.WriteString(command.description)
			result.WriteString("\n")
		}
		return result.String()
	}
}
