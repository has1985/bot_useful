package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/has1985/bot_useful/lib/e"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(chatID int, userName string) (*Page, error)
	Remove(p *Page) error
	IsExist(p *Page) (bool, error)
}

var ErrNoSavedPages = errors.New("no saved page")

type Page struct {
	URL      string
	UserName string
	ID       int
}

func (p Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}
	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
