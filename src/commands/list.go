package commands

import "gopass/data"

func ListCommandHandler(storage *data.Storage, arguments []string) string {
	switch len(arguments) {
	case 0:
		return storage.GetSources()
	case 1:
		return storage.GetCredentials(arguments[0])
	case 2:
		return storage.GetPassword(arguments[0], arguments[1])
	default:
		return getArgumentsLengthMaximumMessage(len(arguments), 2)
	}
}
