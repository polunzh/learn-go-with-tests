package main

const (
	ErrorNotFound = DictionaryErr("unknown word")
	ErrWordExists = DictionaryErr("the word already exists")
	ErrWordDoesNotExist = DictionaryErr("cannt update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (dict Dictionary) Search(word string) (string, error) {

	value, ok := dict[word]

	if !ok {
		return "", ErrorNotFound
	}

	return value, nil
}

func (dict Dictionary) Add(word, value string) error {
	_, err := dict.Search(word)

	switch err {
	case ErrorNotFound:
		dict[word] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Update(word, value string) error {
	_, err := dict.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrWordDoesNotExist
	case nil:
		dict[word] = value
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Delete(word string) error {
	_, err := dict.Search(word)

	if err != nil {
		return err
	}

	delete(dict, word)

	return nil
}
