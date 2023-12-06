package maps

const (
	ErrWordNotFound        = DictionaryErr("word not in dictionary")
	ErrWordExists          = DictionaryErr("word already exists in dictionary")
	ErrNoWordFoundToUpdate = DictionaryErr("word needs to already exist in order to update")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

// Add should not modify existing definitions for a word, it should only add new words
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update should only update a word if it already exists
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrNoWordFoundToUpdate
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}
