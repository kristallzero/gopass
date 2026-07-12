package commands

import (
	"gopass/storage"
	"strings"
)

func HelpCommandHandler(commands []command) func(*storage.Storage, []string) string {
	return func(_ *storage.Storage, _ []string) string {
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
