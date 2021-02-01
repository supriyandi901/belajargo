//variable


/* TYPE DATA:
bool

string

int  int8  int16  int32  int64  // use int unless you want a specific size
uint uint8 uint16 uint32 uint64 uintptr // ditto, use uint

byte // alias for uint8

rune // alias for int32
     // represents a Unicode char

float32 float64

complex64 complex128

*/



package main
import "fmt"

var a, b, c = 10, 20, "tikus"

var (
	s = "aku"
	p = "mencari"
	o = "janda"
	k = 5
)

//constant variable
const ari = "ira"


const (
	wHatever = "whatever"
	one = 1
	janda = "enak"
)



func main() {//variable tanpa deklarasi
	samInt, samBool, samStr := 30, true, "anjani";
	fmt.Println(samInt, samStr, samBool, a, b, c, s, p, o, k)
	fmt.Println(ari, wHatever, one, janda)
}