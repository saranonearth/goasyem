# goasyem
<img src="https://img.shields.io/github/go-mod/go-version/emitter-io/emitter" alt="Go Version"> [![Go Report Card](https://goreportcard.com/badge/github.com/saranonearth/goasynem)](https://goreportcard.com/report/github.com/saranonearth/goasynem) ![example workflow](https://github.com/saranonearth/goasynem/actions/workflows/test.yml/badge.svg)

An idiomatic async emitter for go✨

## Quick Start
```
go get https://github.com/saranonearth/goasynem
```

## Usage Example

```go
 func main() {
  e := &goasynem.Goasynem{}
  
  e.Subscribe("event", func(c interface{}) error {
      fmt.Println(c)
      return nil
    })
    
  e.Emit("event", fmt.Sprintf("Done %d", 1))
 }

```
### Sync Usage

```go
 func main() {
  e := &goasynem.Goasynem{}
  
  e.Subscribe("event", func(c interface{}) error {
      fmt.Println(c)
      return nil
    })
    
  errCh := e.Emit("event", fmt.Sprintf("Done %d", 1))
  err := <-errCh
  if err != nil {
    fmt.Println("Error found", err)
  }
 }

```
## Building and Testing
 Once you have this installed, simply go get this repository and run the following commands to download the package.
```shell
go get -u github.com/saranonearth/goasynem
```
If you want to run the tests, simply run go test command as demonstrated below.

```shell
make test
```

## Contributing
If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.

## License
This is free software under the terms of the MIT license 

