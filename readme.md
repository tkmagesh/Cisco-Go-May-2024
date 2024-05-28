# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
- Commence      : 09:00 AM
- Tea Break     : 10:30 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hour)
- Tea Break     : 03:00 PM (20 mins)
- Wind up       : 05:00 PM

## Methodology
- No powerpoints
- Code & Discussion
- No dedicated time for Q&A

## Repository
[Git](https://github.com/tkmagesh/cisco-go-may-2024)

## Software Requirements
- Go tools (https://go.dev/dl)
    - Verification
        ```go
        go version
        ```
- [Visual Studio Code](https://code.visualstudio.com)
- [Go extensions for VSCode](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- Any Git Client

## Why Go?
- Simplicity
    - ONLY 25 keywords
    - Not(s)
        - No access modifiers (no public/private etc)
        - No classes (ONLY structs)
        - No inheritance (ONLY composition)
        - No reference types (everything is a value)
        - No exceptions (ONLY errors)
        - No try..catch..finally 
        - No implicit type conversion
- Performance
    - Compiled to target platform
    - Tooling support for cross-compilation
    - On par with C++
- Managed (memory) Environment
    - Builtin Garbage Collector
- Concurrency Support
    - Uses lightweight threads called Goroutines (alernatives to OS Threads)
    - Goroutines are extremely cheap (~4kb vs ~2MB for OS Threads)
    - Builtin scheduler to schedule the goroutines
    - Language support for concurrency
        - go keyword, channel (data type), channel operator ( <- ), for..range & select..case language constructs
    - APIs support (standard library)
        - "sync" package
        - "sync/atomic" packages

## Compile & Execute
### Compile
```shell
go build [file_name.go]

go build -o [binary_name] [file_name.go]
```

### Compile & Execute
```shell
go run [file_name.go]
```

### To list all the environment variables
```shell
go env
```

### To list a few environment variables
```shell
go env [var_1] [var_2]
```

### Env variables for cross compilation
- GOOS
- GOARCH

### To get the list of supported platforms (for compilation)
```shell
go tool dist list
```

### To cross compile
```shell
GOOS=[target_os] GOARCH=[target_arch] go build [filename.go]

ex:
GOOS=linux GOARCH=amd64 go build 01-hello-world.go
```

## Data types
- string
- bool
- Integers
    - int8
    - int16
    - int32
    - int64
    - int
- Unsigned Integers
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- Floating point
    - float32
    - float64
- Complex Types
    - complex64 [ real (float32) + imaginary (float32) ]
    - complex128 [ real (float64) + imaginary (float64) ]
- Type aliases
    - byte (alias for uint8)
    - rune (alias for int32, unicode code point)

## Variables
### using 'var' (To declare variables)
```go
    var [var_name] [data_type]
    var [var_1], [var_2] [data_type]
```

### using ':=' (To declare & initialize)
```go
    [var_name] := [data]
```

### Function Scope
- Can use :=
- Cannot have unused variables

### Package Scope
- CANNOT use :=
- Can have unused variables

## IOTA 
    - auto generated constants

## Programming Constructs
### if else
### switch case
### for

## Functions
- More than one return results
- Anonymous functions
- Functions can be assigned to variables
- Higher Order Functions
    - Pass functions as arguments 
    - Return functions as return values
