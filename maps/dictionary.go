package maps

import "errors"

type Dictionary map[string]string

var ErrKeyNotFound = errors.New("could not find the word you looking for")

func (d Dictionary) Search(key string) (string, error) {
    definition, ok := d[key]

	if !ok {
		return "", ErrKeyNotFound
	}

	return definition, nil
}
