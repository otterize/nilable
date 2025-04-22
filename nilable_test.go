package nilable

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullable_IsNull_MarshalJSON(t *testing.T) {
	nullableStr := Nilable[string]{}
	marshalled, err := json.Marshal(nullableStr)
	assert.NoError(t, err)
	assert.JSONEq(t, "null", string(marshalled))
}

func TestNullable_IsNull_FromNil_MarshalJSON(t *testing.T) {
	var nullStr *string = nil
	var nullableStr Nilable[string] = FromPtr(nullStr)
	marshalled, err := json.Marshal(nullableStr)
	assert.NoError(t, err)
	assert.JSONEq(t, "null", string(marshalled))
}

func TestNullable_IsNotNull_MarshalJSON(t *testing.T) {
	nullableStr := From("test")
	marshalled, err := json.Marshal(nullableStr)
	assert.NoError(t, err)
	assert.JSONEq(t, "\"test\"", string(marshalled))
}

func TestNullable_IsNull_UnmarshalJSON(t *testing.T) {
	var out Nilable[string]
	err := json.Unmarshal([]byte("null"), &out)
	assert.NoError(t, err)
	assert.False(t, out.Set)
}

func TestNullable_IsNotNull_UnmarshalJSON(t *testing.T) {
	var out Nilable[string]
	err := json.Unmarshal([]byte("\"test\""), &out)
	assert.NoError(t, err)
	assert.True(t, out.Set)
	assert.Equal(t, "test", out.Item)
}

func TestNullable_Valuer(t *testing.T) {
	nullableStr := From("test")
	val, err := nullableStr.Value()
	assert.NoError(t, err)
	assert.Equal(t, "test", val)
}

func TestNullable_Scanner(t *testing.T) {
	var out Nilable[string]
	err := out.Scan(nil)
	assert.NoError(t, err)
	assert.False(t, out.Set)

	err = out.Scan("test")
	assert.NoError(t, err)
	assert.True(t, out.Set)
	assert.Equal(t, "test", out.Item)
}

func TestNullable_ValuerScannerBackAndForth(t *testing.T) {
	nullableStr := From("test")
	val, err := nullableStr.Value()
	assert.NoError(t, err)

	var out Nilable[string]
	err = out.Scan(val)
	assert.NoError(t, err)
	assert.True(t, out.Set)
	assert.Equal(t, "test", nullableStr.Item)
}
