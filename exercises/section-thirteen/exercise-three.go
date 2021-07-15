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
	veh := vehicle{
		doors: 5,
		color: "red",
	}

	taco := truck{
		vehicle: vehicle{
			doors: 4,
			color: "silver",
		},
		fourWheel: true,
	}

	wrx := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(veh, taco, wrx)
	fmt.Println(veh.doors, taco.doors, wrx.doors)
}
