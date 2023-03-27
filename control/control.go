package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	// sleep for 1 min
	robotgo.MouseSleep = 1 * 1000

	// Start moving cursor
	fmt.Println("Moving mouse...")

	for i := 1; i <= 5000; i++ {
		x := 1 * i
		y := 1 * i

		if i%2 == 0 {
			robotgo.MoveRelative(-x, -y)
			fmt.Println(fmt.Sprintf("%v, %v", -x, -y))
		} else {
			robotgo.MoveRelative(x, y)
			fmt.Println(fmt.Sprintf("%v, %v", x, y))
		}
	}
}
