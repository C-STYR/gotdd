package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find key in dictionary")
	ErrKeyDoesNotExist  = DictionaryErr("could not locate key to be updated")
	ErrKeyAlreadyExists = DictionaryErr("key already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	_, ok := d[key]
	if ok {
		return ErrKeyAlreadyExists
	}
	d[key] = val
	return nil
}

func (d Dictionary) Update(key, val string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrKeyDoesNotExist
	case nil:
		d[key] = val
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
