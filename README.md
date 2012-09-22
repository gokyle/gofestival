## gofestival

gofestival is a Go interface to festival's text-to-speech engine.

usage:
```go
    import "bitbucket.org/kisom/gofestival"

    ...
    err := festival.Speak("Hello, world.")
    if err != nil {
            fmt.Println(err)
    }
```
