Exercise: rpncalc

Instructions
Write a program that takes a string which contains an equation written in Reverse Polish Notation (RPN) as its first argument, evaluates the equation, and prints the result followed by a newline ('\n').

Operands and operators are separated by spaces. The supported operators are: +, -, *, /, %.

If the string is invalid or if there is not exactly one argument, print "Error" followed by a newline. Extra spaces are allowed.

Examples
$ go run . "1 2 * 3 * 4 +" | cat -e
10$
$ go run . "1 2 3 4 +" | cat -e
Error$
$ go run . | cat -e
Error$
$ go run . "     1      3 * 2 -" | cat -e
1$
$ go run . "     1      3 * ksd 2 -" | cat -e
Error$
