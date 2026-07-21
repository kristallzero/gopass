package commands

import (
	"gopass/data"
)

func DeleteCommandHandler(storage *data.Storage, arguments []string) string {
	if err := checkArgumentsLengthRange(arguments, 1, 2); len(err) != 0 {
		return err
	}
	if !storage.SourceExists(arguments[0]) {
		return "error: source hasn't been found"
	}
	if len(arguments) == 1 {
		if !storage.IsOneCredential(arguments[0]) {
			return "error: there are more than one crendetial in source, login can't be omitted"
		}
		return storage.DeleteOneCredential(arguments[0])
	}
	return storage.DeleteCredential(arguments[0], arguments[1])
}
