package maps

import "errors"

type Dictionary map[string]string

var ErrWordNotFound = errors.New("word not in dictionary")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
