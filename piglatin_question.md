Exercise: piglatin

Instructions
Write a program that transforms a string passed as argument in its Pig Latin version.

Rules:
- If a word begins with a vowel, add "ay" to the end.
- If it begins with a consonant, take all consonants before the first vowel and move them to the end of the word, then add "ay".
- Only the Latin vowels will be considered as vowels (a e i o u). Handle case-insensitively.
- If the word has no vowels, print "No vowels".
- If the number of arguments is different from one, the program prints nothing.

Examples
$ go run . pig | cat -e
igpay$
$ go run . Is | cat -e
Isay$
$ go run . crunch | cat -e
unchcray$
$ go run . crnch | cat -e
No vowels$
