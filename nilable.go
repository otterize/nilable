package nullable

import (
	"bytes"
	"encoding/json"
	"reflect"
)

type Nilable[T any] struct {
	Item T
	Set  bool
}

func From[T any](item T) Nilable[T] {
	reflectVal := reflect.ValueOf(item)
	if reflectVal.Kind() == reflect.Ptr && reflectVal.IsNil() {
		return Nilable[T]{Set: false}
	}
	return Nilable[T]{Item: item, Set: true}
}

func (n Nilable[T]) MarshalJSON() ([]byte, error) {
	if !n.Set {
		return []byte("null"), nil
	}
	return json.Marshal(n.Item)
}

func (n *Nilable[T]) UnmarshalJSON(data []byte) error {
	if data == nil || bytes.Equal(data, []byte("null")) {
		n.Set = false
		return nil
	}

	err := json.Unmarshal(data, &n.Item)
	if err != nil {
		return err
	}
	n.Set = true
	return nil
}
