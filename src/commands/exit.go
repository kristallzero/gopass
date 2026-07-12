package commands

import "gopass/storage"

func ExitCommandHandler(storage *storage.Storage, _ []string) string {
	storage.SaveCredentials()
	return "exit"
}
