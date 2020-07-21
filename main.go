package goarea

import "math"

//Pi é uma proporção entre o perímetro e o diâmetro
const Pi = 3.1416

func Circ(r float64) float64 {
	return math.Pow(r, 2) * Pi
}

func Rect(base, altura float64) float64 {
	return base * altura
}

func esfera(r float64) float64 {
	return 4 * Pi * math.Pow(r, 2)
}

//Não é visível
func _Triang(base, altura float64) float64 {
	return (base * altura) / 2
}
