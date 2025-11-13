Question: rostring — rotate the first word of the string to the end; words separated by spaces in output. If number of arguments != 1, print a newline.

Rules:
- A word is a sequence of alphanumerical characters (for this exercise we treat any non-space "token" as a word).
- Output words separated by a single space and followed by a newline.
- If the argument contains only spaces or is empty, print a newline.

Usage:
$ go run . "abc   " | cat -e
abc$
$ go run . "Let there     be light"
there be light Let

Answer file: rostring.go
