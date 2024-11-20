package maps

type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

var (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("the key you are trying to add is already existed")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value

	case nil:
		return ErrWordExists

	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, newVal string) {
	d[key] = newVal
}