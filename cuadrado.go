package figuras

type Cuadrado struct{
	Lado float64 // Lado publico
}

func (cua *Cuadrado) area() float64 {
	return	cua.Lado * cua.Lado
}

func (cua *Cuadrado) perimetro() float64 {
	return  4 * cua.Lado 
}