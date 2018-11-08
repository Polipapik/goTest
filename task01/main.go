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

// Bro - Your homie
type Bro struct {
	Guy            // Girl can't be a Real brro..
	Description    string
	Cash           int
	BestJokesTheme string
	HasGirl        bool
	Virtual        bool
}

// GrowOld function (постареть) ((пусть они стареют так))
func (g *Guy) GrowOld() {
	g.Age++
	g.Height--
	g.Weight++
}

// BodyMassIndex - Индекс массы тела
func BodyMassIndex(height, weight int) float64 {
	return float64(weight*10000) / float64(height*height)
}

func main() {

	vasya := Guy{
		Name:       "Vasya",
		Age:        15,
		Height:     150,
		Weight:     150,
		HasHousing: true,
	}

	fmt.Println(vasya.Name, "GROW OLD")
	fmt.Println("Before:", vasya)
	vasya.GrowOld()
	fmt.Println("After:", vasya)
	fmt.Println("\nBody mass index -", BodyMassIndex(vasya.Height, vasya.Weight))

	vasyaBro := Bro{
		Guy:         vasya,
		Description: "brat za brata",
	}
	fmt.Println("\nIf Vasya is bro:", vasyaBro)

	fmt.Scanln()
}
