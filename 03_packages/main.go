package main

import (
	"fmt"
	"math"

	"github.com/cesarlugoe/crashcourse/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(5.4))
	fmt.Println(math.Ceil(5.4))
	fmt.Println(math.Sqrt(4))
	fmt.Println(strutil.Reverse("aloh"))
}
