package commands

import "gopass/data"

func ExitCommandHandler(storage *data.Storage, _ []string) string {
	storage.SaveCredentials()
	return "exit"
}
