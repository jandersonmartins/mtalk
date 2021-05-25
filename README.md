# Mtalk

Simple way to run os commands through a TCP connection

## Install

```sh
$ go get github.com/jandersonmartins/mtalk
```

## Usage

```go
package main

import (
	"github.com/jandersonmartins/mtalk"
)

func main() {
	mtalk.Listen(8081)
}
```