package telegram

import "github.com/has1985/bot_useful/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client) {

}
