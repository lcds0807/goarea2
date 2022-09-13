package goarea2

import "math"

// Pi é uma proporção numérica definida pela relação entgre
// o perimetro de uma circunferencia e seu diametro

const Pi = 3.1416

// Circ é resp por calc a area da circunf
func Circ(raio float64) float64 {
	return math.Pow(raio, 2)
}

// Rect é resp por calc a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
