package dictionary

const (
	NotFoundError   = DictionaryError("value not found in Dictionary")
	WordExistsError = DictionaryError("word already exists in Dictionary")
  WordDoesNotExistError = DictionaryError("word does not exist in Dictionary to update")
)

type DictionaryError string
type Dictionary map[string]string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {

	var err error = nil

	match, ok := d[key]

	if !ok {
		err = NotFoundError
	}
	return match, err
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case NotFoundError:
		d[key] = value
	case nil:
		return WordExistsError
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
  _, err := d.Search(key)

switch err {
  case NotFoundError:
    return WordDoesNotExistError
  case nil:
    d[key] = value
  default:
    return err
}
  return nil
}
