package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/has1985/bot_useful/lib/e"
	"github.com/has1985/bot_useful/lib/storage"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	basePath string
}

const (
	defaultPerm = 0755
)

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {

	defer func() { err = e.WrapIfErr("can't save page", err) }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.Mkdir(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}
	return nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = e.WrapIfErr("can't pick random page", err) }()

	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New(storage.ErrNoSavedPages.Error())
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))
}

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)
	if err != nil {
		return e.Wrap("can't remove file", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path); err != nil {
		return e.Wrap(fmt.Sprintf("can't remove file %s", path), err)
	}
	return nil
}

func (s Storage) IsExist(p *storage.Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, e.Wrap("can't check if file exist", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("can't check if file %s exist", path)
		return false, e.Wrap(msg, err)
	}
	return true, nil
}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("can't open file", err)
	}

	defer func() { _ = file.Close() }()

	var p storage.Page

	if err := gob.NewDecoder(file).Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
