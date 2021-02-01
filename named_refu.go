//named return function

package main
import "fmt"

func dasar(x, y int) (xtambah2 int, ytambah2 int) {
	xtambah2 = x+2
	ytambah2 = y+2

	return xtambah2, ytambah2
}
func init() {
	fmt.Println("executing init function")
}

func main() {
	fmt.Println(dasar(20,30))
}