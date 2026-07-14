package commands

import (
	"gopass/data"
)

func EditCommandHandler(storage *data.Storage, arguments []string) string {
	if err := checkArgumentsLengthRange(arguments, 3, 4); len(err) != 0 {
		return err
	}
	if !storage.SourceExists(arguments[0]) {
		return "error: source hasn't been found"
	}
	if len(arguments) == 3 {
		if !storage.IsOneCredential(arguments[0]) {
			return "error: there are more than one crendetial in source, login can't be omitted"
		}
		return storage.UpdateOneCredential(arguments[0], arguments[1], arguments[2])
	}
	return storage.UpdateCredential(arguments[0], arguments[1], arguments[2], arguments[3])
}
