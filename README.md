# ANSI Escape Sequences [![GoDoc](https://godoc.org/github.com/alan-b-lima/ansi-escape-sequences?status.svg)](https://pkg.go.dev/github.com/alan-b-lima/ansi-escape-sequences)

ANSI Escape Sequences is a minimal Go library to abstract ANSI escape sequences. It also offers functionality to enable and disable escape sequences on Windows 10 v1511 and later.

> [!WARNING]
> ANSI Escape Sequences is still under development, and is subject to breaking changes.

# Pen

Pens are a higher-level abstraction over ANSI escape sequences. They allow you to create reusable styles that can be applied to text. The snippet below will print "Hello, World!" in bold, italic, red text on a black background:

```go
package main

import (
    "fmt"
    "os"

    "github.com/alan-b-lima/ansi-escape-sequences"
)

func main() {
    pen := &ansi.Pen{Writer: os.Stdout}
    
    pen.Bold()
    pen.Italic()
    pen.FGColor(ansi.RGB{R: 255})
    pen.BGColor(ansi.RGB{})

    fmt.Fprintln(pen, "Hello, World!")
}
```


Pens can also be used to fire cursor movements and screen manipulations on the underlying writer.

```go
package main

import (
    "fmt"
    "os"

    "github.com/alan-b-lima/ansi-escape-sequences"
)

func main() {
    pen := &ansi.Pen{Writer: os.Stdout}
    
    pen.EraseScreen()
    pen.MoveCursor(0, 0)

    fmt.Fprintln(pen, "Hello, World!")
}
```