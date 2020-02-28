package metadata

func NewMemory() Store {
	return memStore{m: map[string]interface{}{}}
}

type memStore struct {
	m map[string]interface{}
}

func (ms memStore) Read(keys ...string) (string, error) {
	if len(keys) == 0 {
		return "", ErrNoKeys
	}
	m, err := ms.getMap(keys[:len(keys)-1])
	if err != nil {
		return "", err
	}
	switch t := m[keys[len(keys)-1]].(type) {
	case map[string]interface{}:
		return "", ErrNotValue
	default:
		return primToString(t)
	}
}

func (ms memStore) getMap(keys []string) (map[string]interface{}, error) {
	m := ms.m
	for _, key := range keys {
		switch t := m[key].(type) {
		case map[string]interface{}:
			m = t
		default:
			return nil, ErrNotValue
		}
	}
	return m, nil
}


func (ms memStore) ReadAll() (map[string]interface{}, error) {
	return copyMap(ms.m), nil
}

func copyMap(m map[string]interface{}) map[string]interface{} {
	out := map[string]interface{}{}
	for k, v := range m {
		if vm, ok := v.(map[string]interface{}); ok {
			out[k] = copyMap(vm)
		} else {
			out[k] = v
		}
	}
	return out
}

func (ms memStore) Delete(keys ...string) error {
	if len(keys) == 0 {
		return ErrNoKeys
	}
	m, err := ms.getMap(keys[:len(keys)-1])
	if err != nil {
		return err
	}
	delete(m, keys[len(keys)-1])
	return nil
}

func (ms memStore) DeleteAll() error {
	ms.m = map[string]interface{}{}
	return nil
}

func (ms memStore) Write(value string, keys ...string) error {
	if len(keys) == 0 {
		return ErrNoKeys
	}
	m, err := ms.getMap(keys[:len(keys)-1])
	if err != nil {
		return err
	}
	m[keys[len(keys)-1]] = value
	return nil
}

func (ms memStore) WriteAll(metadata map[string]interface{}) error {
	ms.m = copyMap(metadata)
	return nil
}
