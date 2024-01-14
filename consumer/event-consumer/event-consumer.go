package event_consumer

import (
	"github.com/has1985/bot_useful/events"
	"log"
	"time"
)

type Consumer struct {
	fetcer    events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) *Consumer {
	return &Consumer{
		fetcer:    fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start() error {
	for {
		gotEvents, err := c.fetcer.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())
			//retry
			continue
		}
		////////////////////////////////////////////////////////////
		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}
	}
}

func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new events:%s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())
			continue
		}
	}
	return nil
}
