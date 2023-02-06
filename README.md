# nilable
nilable is a tiny Go generics library for making non-`nil`able values nilable, and supports JSON serialization.

You too can stop using pointers when you don't need them!

## From value to JSON
```go
nilableString := nilable.From("hello world")
println(nilableString.Set) // true
println(nilableString.Item) // "hello world"
data, _ := json.Marshal(nilableString)
println(string(data)) // `"hello world"` 
```

## From null JSON to value
```go
var nilJson Nilable[string]
json.Unmarshal([]byte("null"), &nilJson)
println(nilJson.Set) // false
```