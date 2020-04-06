# Go-No-Go [![Go Report Card](https://goreportcard.com/badge/github.com/GrantEthanEckstein/Go-No-Go)](https://goreportcard.com/report/github.com/GrantEthanEckstein/Go-No-Go) [![GoDoc](https://godoc.org/github.com/GrantEthanEckstein/Go-No-Go?status.svg)](https://godoc.org/github.com/GrantEthanEckstein/Go-No-Go)

Example Usage
```go
package main

import (
	"gonogo"
	"fmt"
)

func function1(data []byte) []byte {
	return append(data, []byte(" is")...)
}

func function2(data []byte) []byte {
	return append(data, []byte(" awesome!")...)
}

func main() {
	a := gonogo.NewNegotiation("test")
	a.AddStep("Client", function1)
	a.AddStep("Client", function2)

	r := a.GetRole("Client")

	data := r.Execute([]byte("Grant"))
	fmt.Println(string(data))
}

```
