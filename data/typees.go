package data

import "fmt"

type location string
type distance int

func (origin location) DistanceTo(distination location) distance {
	fmt.Printf(("Origin %v Detsination %v\n"), origin, distination)
	return 10
}

func locationTest() {
	nyx := location("Nyx")
	nyx.DistanceTo(location("-23, -44"))
}
