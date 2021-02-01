//raw string
/*
Go has two types of strings:
-> Interpreted strings: The typical string type created with ". Can contain anything except new line and unescaped ".
-> Raw strings: Encoded between "`" (backticks) can contain new lines and other artifacts.

*/

package main
import "fmt"

func main() {
	rawstr :=

        `First line 

some new lines

more new lines

"double quotes"
    `

    fmt.Print(rawstr)
}