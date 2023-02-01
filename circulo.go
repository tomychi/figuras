package figuras

import "math"

type Circulo struct {
	Radio float64 // Radio en mayuscula para que sea publico
}

func(cir *Circulo) area() float64 {
	return math.Pi * cir.Radio * cir.Radio
}

func(cir *Circulo) perimetro() float64 {
	return math.Pi * 2 * cir.Radio
}