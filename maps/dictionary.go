package MapsTest

// import "errors"

type DictionaryErr string

const (
	errorNotFound = DictionaryErr("could not find the word you were looking for")
	errorWordExits = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// var errorNotFound = errors.New("could not find the word you were looking for")
// var errorWordExits = errors.New("cannot add word because it already exists")
func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", errorNotFound
	}

	return value, nil
}

func (d Dictionary) Add(word, value string) error {
	// d[word] = value
	// return nil
	_, err := d.Search(word)

	switch err {
	case errorNotFound:
		d[word] = value
	case nil:
		return errorWordExits
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, value string) error {
	_, err := d.Search(word)
	
	switch err {
	case errorNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case errorNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}