package telegram

const msgHelp = `I can save and keep you pages. Also I can offer you them to read.

In order to save a page, just send me a link to it.

In order to get a random page from your list, send me /rnd command.
Caution! After thet, this page will be removed from your list!`

const msgHello = "Hi there! \n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command."
	msgNoSavedPages   = "You have no saved pages."
	msgSaved          = "Saved."
	msgAlreadyExist   = "You have already have this page in your list."
)
