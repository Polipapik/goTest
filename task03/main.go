package main

import "fmt"

// Body comment
type Body interface {
	GrowOld()
	Description() string
}

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
	Guy
	Quote          string
	Cash           int
	BestJokesTheme string
	HasGirl        bool
	Virtual        bool
}

// BodyMassIndex comment
func BodyMassIndex(height, weight int) float64 {
	return float64(weight*10000) / float64(height*height)
}

// GrowOld function
func (g *Guy) GrowOld() {
	g.Age++
	g.Height--
	g.Weight++
}

// Description - описание в виде строки, + бодимасиндекс
func (g Guy) Description() string {
	result := fmt.Sprintf("Name: %s, Age: %d, Height: %d, Weight: %d, HasHousing: %t, BMI: %f",
		g.Name, g.Age, g.Height, g.Weight, g.HasHousing,
		BodyMassIndex(g.Height, g.Weight))
	return result
}

// GrowOld бро стареют по-другому)
func (g *Bro) GrowOld() {
	g.Age++
	g.Height += 5
	g.Weight--
}

// Description (все поля не выводил)
func (g Bro) Description() string {
	result := fmt.Sprintf("BRO Name: %s, QUOTE: %s, Cash: %d, Age: %d, Height: %d, Weight: %d, HasHousing: %t, BMI: %f",
		g.Name, g.Quote, g.Cash, g.Age, g.Height, g.Weight, g.HasHousing,
		BodyMassIndex(g.Height, g.Weight))
	return result
}

func main() {
	var bodies [4]Body

	bodies[0] = &Guy{
		Name:       "Vasya",
		Age:        15,
		Height:     150,
		Weight:     150,
		HasHousing: true,
	}
	bodies[1] = &Guy{
		Name:       "Kirill",
		Age:        18,
		Height:     176,
		Weight:     92,
		HasHousing: false, // думал, писать или нет
	}
	bodies[2] = &Bro{
		Guy: Guy{
			Name:       "Kostya",
			Age:        25,
			Height:     220,
			Weight:     50,
			HasHousing: true,
		},
		Quote: "kek",
	}
	bodies[3] = &Guy{
		Name:       "Petya",
		Age:        20,
		Height:     180,
		Weight:     70,
		HasHousing: false, // думал, писать или нет
	}

	for i := 0; i < 4; i++ {
		fmt.Println(i, "|", bodies[i].Description())
	}

	fmt.Print("\nAfter 1 year:\n\n")

	for i := 0; i < 4; i++ {
		bodies[i].GrowOld()
		fmt.Println(i, "|", bodies[i].Description())
	}
	fmt.Scanln()
}
