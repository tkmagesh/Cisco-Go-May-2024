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
        - No implicity type conversion
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

