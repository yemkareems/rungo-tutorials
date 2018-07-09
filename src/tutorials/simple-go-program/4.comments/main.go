// this is main package
package main

/*
Import multiple package but use only fmt package.
To avoid compilation error, use _ (blank identifier) as alias
*/
import (
	"fmt"
	_ "math"
	_ "net/http"
)

func main() {
	fmt.Println("Hello World!")
}
