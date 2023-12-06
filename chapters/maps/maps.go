package maps

import (
	"errors"
)

type Dictionary map[string]string

var ErrWordNotFound = errors.New("word not in dictionary")
var ErrWordExists = errors.New("word already exists in dictionary")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

// Add should not modify existing definitions for a word, it should only add new words
func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]

	if !ok {
		d[word] = definition
		return nil
	}
	return ErrWordExists
}
