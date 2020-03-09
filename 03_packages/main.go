package main

import (
	"fmt"
	"math"

	"github.com/MikaelOttosson/golang-cc/03_packages/strutil"
)

func main() {
	fmt.Println("Floor", math.Floor(2.5))
	fmt.Println("Ceil", math.Ceil(2.5))
	fmt.Println("Ceil", math.Sqrt(24))

	fmt.Println(strutil.Reverse("ekiM ,olleH"))
}
