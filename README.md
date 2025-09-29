# ANSI Escape Sequences [![GoDoc](https://godoc.org/github.com/alan-b-lima/ansi-escape-sequences?status.svg)](https://pkg.go.dev/github.com/alan-b-lima/ansi-escape-sequences)

ANSI Escape Sequences is a minimal Go library to abstract ANSI escape sequences. It also offers functionality to enable and disable escape sequences on Windows 10 v1511 and later.

> [!WARNING]
> ANSI Escape Sequences is still under development, and is subject to breaking changes.

## Summary

This package focuses on abstracting ANSI escape sequences in a friendly interface. There are three main ways to use this package, through pens, builders and top-level functions.

### Pens


### Builders

Builders take heavy inspirations from Go's own strings.Builder type. It adds flexibility to append escape sequences supported by this package, and it may be, at the end of accumulation, flushed to some stream, usually os.Stdout. It implements the io.Writer interface, therefore it can be used to alongside fmt.Fprint, fmt.Fprintf and fmt.Fprintln functions.

### Top-Level Functions

Top-level functions related to escape sequences, in this package, always returns a string that does, if printed to a complient virtual terminal, what it is intended to. See the [final chapter](#about-ansi-escape-sequences) for more info on the supported sequences.

## Enabling and Disabling Virtual Terminal Processing

On some Windows' terminals, like CMD, the processing of ANSI escape sequences is not done by default, and you'd end up with a bunch of weird characters by printing such sequences. To solve that, for Windows builds, two functions are provided to enable and disable such processing, this functions are also present for other builds, but they do nothing. You may have the following piece of code at the start of your program:

```go
if err := ansi.EnableVirtualTerminal(os.Stdout.Fd()); err != nil {
	panic(err)
}
defer ansi.DisableVirtualTerminal(os.Stdout.Fd())
```

Currently, there are no mock versions to handle the case where the provided standart output cannot process escape sequences.

## About ANSI Escape Sequences

This is a comprehensive list of all the ANSI escape sequences supported by this package.

### Graphic Style Settings

Select Graphic Rendition `ESC` `[` `<attr>` {`;` `<attr>`} `m`

* `<attr>`:
    * bold `1`
    * italic `3`
    * underline `4`
    * strike `9`
    * normal intensity `22`
    * not italic `23`
    * not underline `24`
    * not crossed `29`
    * set foreground (24 bits) `38` `;` `2` `;` `<red>` `;` `<green>` `;` `<blue>`:
        * `<red>`: a decimal number in the range 0-255
        * `<green>`: a decimal number in the range 0-255
        * `<blue>`: a decimal number in the range 0-255
    * default foreground `39`
    * set background (24 bits) `49` `;` `2` `;` `<red>` `;` `<green>` `;` `<blue>`:
        * `<red>`: a decimal number in the range 0-255
        * `<green>`: a decimal number in the range 0-255
        * `<blue>`: a decimal number in the range 0-255
    * default background `49`

### Hyperlink

HyperLink `ESC` `8` `;` `;` `<link>` `ESC` `\` `<text>` `ESC` `8` `8` `;` `;` `ESC` `\`

* `<link>`: a absolute URI, should contain scheme
* `<text>`: any string, may be styled

### Cursor, Screen and Devide Control

Move Cursor Up `ESC` `[` `<n>` `A`
* move the cursor up by `n` rows

Move Cursor Down  `ESC` `[` `<n>` `B`
* move the cursor down by `n` rows

Move Cursor Left  `ESC` `[` `<n>` `D`
* move the cursor left by `n` columns

Move Cursor Right `ESC` `[` `<n>` `C`
* move the cursor right by `n` columns

Move Cursor To `ESC` `[` `<r>` `;` `<c>` `H`
* move the cursor to the `r`-th row and `c`-th columns, both are 1-indexed

Scroll Up   `ESC` `[` `<n>` `S`
* scroll up by `n` columns

Scroll Down `ESC` `[` `<n>` `T`
* scroll down by `n` columns

Erase Screen `ESC` `[` `2` `J`
* erases the entire screen

Erase Line `ESC` `[` `2` `K`
* erases the current line

Style Cursor `ESC` `[` `<style>` ` ` `q`
* `<style>`:
    * blinking block `0`, `1` or none
    * steady block `2`
    * blinking underline `3`
    * steady underline `4`

Show Cursor `ESC` `[` `?` `25` `h`
* show the cursor

Hide Cursor `ESC` `[` `?` `25` `l`
* hide the cursor

EnterAlt = `ESC` `[` `?` `1049` `h`
* enter alternate screen buffer

LeaveAlt = `ESC` `[` `?` `1049` `l`
* leave alternate screen buffer

Enter Bracketed Paste Mode `ESC` `[` `?` `2004` `h`
* enter bracketed paste mode, pasted text will be surrounded by `ESC` `[` `200` `~` and `ESC` `[` `201` `~`

Leave Bracketed Paste Mode `ESC` `[` `?` `2004` `l`
* leave bracketed paste mode