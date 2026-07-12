package commands

import "gopass/data"

func ListCommandHandler(storage *data.Storage, _ []string) string {
	return storage.String()
}
