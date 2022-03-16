package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "silver",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors, t.color, t.fourWheel)
	fmt.Println(s.doors, s.color, s.luxury)
}
