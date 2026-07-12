package commands

import "gopass/storage"

func ListCommandHandler(storage *storage.Storage, _ []string) string {
	return storage.String()
}
