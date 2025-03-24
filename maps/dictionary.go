package maps

type Dictionary map[string]string

const (
	ErrKeyNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
    definition, ok := d[key]

	if !ok {
		return "", ErrKeyNotFound
	}

	return definition, nil
}


func (d Dictionary) Add(key string, definition string) error {
    _, err := d.Search(key)

    switch err {
        case ErrKeyNotFound:
            d[key] = definition
        case nil:
            return ErrWordExists
        default:
            return err
    }
    return nil
}


func (d Dictionary) Update(key string, definition string) error {
    _, err := d.Search(key)

    switch err {
        case ErrKeyNotFound:
            return ErrWordDoesNotExist
        case nil:
            d[key] = definition
        default:
            return err
    }
    return nil
}


func (d Dictionary) Delete(key string) error {
    _, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}

