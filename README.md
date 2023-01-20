# 🦫 my Go playground

My Go playground while learning Go. 👩‍🚒  
so be cautious that the codes is not sophisticated and may not follow the best practice.

Run playground simply by `go test`: 🚀

```sh
go test ./...
```

## 🛝 Playground List

- [`nil_playground.go`](internal/nil_playground.go) | [`nil_playground_test.go`](internal/nil_playground_test.go)
  - check the behavior of `nil` slice/map
- [`string_playground.go`](internal/string_playground.go) | [`string_playground_test.go`](internal/string_playground_test.go)
  - check the difference between `len` and `rune count` of string
  - conversion of
    - `string` <=> `[]rune`
    - `rune` => `string`
    - `string` <=> `[]byte`
    - `byte` => `string`
    - `[]rune` <=> `[]byte`
  - string iteration
    - `for-range` loop iteration return RUNE value with BYTE index, NOT return byte value,  
      make mistakes very easily, need caution!
    - iterate with byte index only of string by `for-range` loop, lose bytes! NONSENSE!!
  - check memory struct of string by `SizeOf`
- [`empty_struct_channel_playground.go`](internal/empty_struct_channel_playground.go) | [`empty_struct_channel_playground_test.go`](internal/empty_struct_channel_playground_test.go)
  - read channel by `for-range` loop, very convenient
  - read channel by `comma-ok` pattern, more tedious but safe too

## 😈 My common functions and go design experiments

- [`commons.go`](internal/commons.go)
  - in my personal opinion,  
    `type voidT = struct{}` and `var void = voidT{}`  
    definitions seems interesting and COOL
