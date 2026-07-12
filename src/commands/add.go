package commands

import (
	"gopass/data"
)

func AddCommandHandler(storage *data.Storage, arguments []string) string {
	if err := checkArgumentsLength(arguments, 3); len(err) != 0 {
		return err
	}
	return storage.AddCredentials(arguments[0], arguments[1], arguments[2])
}
