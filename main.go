package main

import (
	"fmt"
	"github.com/golangatspeed/generictesting/generic"
)

func main() {

	fmt.Println(generic.AddAny("hello ", "world"))
	fmt.Println(generic.AddAny(1, 2))
}
