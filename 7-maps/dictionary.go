package main

// map
// a gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attemps to write to a nil map will cause a runtime panic.

// 1/ never initialize an empty map variable
// var m map[string]string

// 2/ two ways to initialize safely. Both approaches create an empty hash map and point dictionary at it. Which ensures that will never get a runtime panic
// var dictionary = map[string]string{}
// var dictionary = make(map[string]string)

// error constant and error wrapper
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// map lookup can return two values
	// the second value is a boolean which indicates if the key was found successfully
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
