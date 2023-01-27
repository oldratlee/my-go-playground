# 🦫 my Go playground

My Go playground while learning Go. 🧑‍🚒  
so be cautious that the codes is not sophisticated and may not follow the best practice.

Run playground simply by `go test`: 🚀

```sh
go test ./...
```

## 🛝 Playground List

- [`nil_empty_playground.go`](internal/nil_empty_playground.go) | [`nil_empty_playground_test.go`](internal/nil_empty_playground_test.go)
  - check the behavior of `nil`/empty slice
  - check the behavior of `nil`/empty map
    - put new elements into nil map, PANIC!
  - check the behavior of `nil` function/lambda
    - call nil function/lambda, PANIC!
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
- [`goroutine_channel_playground.go`](internal/goroutine_channel_playground.go) | [`goroutine_channel_playground_test.go`](internal/goroutine_channel_playground_test.go)
  - read channel by `for-range` loop, very convenient
  - read channel by `comma-ok` pattern, more tedious but safe too
- [`goroutine_channel_pattern_playground.go`](internal/goroutine_channel_pattern_playground.go) | [`goroutine_channel_pattern_playground_test.go`](internal/goroutine_channel_pattern_playground_test.go)
  - operation(read/write) to multiply channels with `select statement` avoid starvation/deadlock
  - `done channel pattern` aka. a channel only used to publish close event
  - `cancel function pattern` to terminate a goroutine
  - `nilify the case channel pattern` to turning off a case in a select
  - `time after case pattern` to time out code
- [`when_go_runtime_panic_playground.go`](internal/when_go_runtime_panic_playground.go) | [`when_go_runtime_panic_playground_test.go`](internal/when_go_runtime_panic_playground_test.go)
  - `array operation` panic
  - `slice operation` panic
  - `channel operation` panic
  - `function value operation`, call nil function value panic
  - call `value receiver method` with nil point receiver panic
  - call `method of nil interface` panic
  - call `method of interface with nil value` panic
  - `pointer operation`, dereference a nil pointer panic
  - `type assertions` panic

## 😈 My common functions and go design experiments

- [`commons.go`](internal/commons.go)
  - in my personal opinion,  
    `type voidT = struct{}` and `var void = voidT{}`  
    definitions seems interesting and COOL
