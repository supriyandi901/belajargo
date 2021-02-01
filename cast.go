//casting -> need to be explicit

package main
import "fmt"

func main() {
	var a, b = 20, 30
	fmt.Println("need to convert a and b to float32 before devide")
	var div float32 = float32(a) / float32(b)
	fmt.Println("Cast float32 to int")
	var divint = int(div)
	fmt.Println(div, divint)
}