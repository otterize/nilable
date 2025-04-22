package nilable

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
)

type Nilable[T any] struct {
	Item T
	Set  bool
}

// From converts from non-pointer types to Nilable.
func From[T any](item T) Nilable[T] {
	reflectVal := reflect.ValueOf(item)
	if (reflectVal.Kind() == reflect.Ptr || reflectVal.Kind() == reflect.Slice || reflectVal.Kind() == reflect.Map) && reflectVal.IsNil() {
		return Nilable[T]{Set: false}
	}
	return Nilable[T]{Item: item, Set: true}
}

// FromPtr converts from pointer types to Nilable.
func FromPtr[T any](item *T) Nilable[T] {
	if item == nil {
		return Nilable[T]{Set: false}
	}
	return Nilable[T]{Item: *item, Set: true}
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

func (n Nilable[T]) Value() (driver.Value, error) {
	if !n.Set {
		return nil, nil
	}
	return n.Item, nil
}

func (n *Nilable[T]) Scan(value any) error {
	if value == nil {
		n.Set = false
		return nil
	}

	item, ok := value.(T)
	if !ok {
		return errors.New("failed to scan value into Nilable")
	}
	n.Item = item
	n.Set = true
	return nil
}
