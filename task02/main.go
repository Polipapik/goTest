package main

import "fmt"

// Guy - Just a guy
type Guy struct {
	Name       string
	Age        int
	Height     int // cm
	Weight     int // kg
	HasHousing bool
}

func main() {
	var guys [4]Guy

	guys[0] = Guy{
		Name:       "Vasya",
		Age:        15,
		Height:     150,
		Weight:     150,
		HasHousing: true,
	}
	guys[1] = Guy{
		Name:       "Kirill",
		Age:        18,
		Height:     176,
		Weight:     92,
		HasHousing: false, // думал, писать или нет
	}
	guys[2] = Guy{
		Name:       "Kostya",
		Age:        25,
		Height:     220,
		Weight:     50,
		HasHousing: true,
	}
	guys[3] = Guy{
		Name:       "Petya",
		Age:        20,
		Height:     180,
		Weight:     70,
		HasHousing: false, // думал, писать или нет
	}

	for i := 0; i < 4; i++ {
		fmt.Println(i, "|", guys[i])
	}

	fmt.Scanln()
}
