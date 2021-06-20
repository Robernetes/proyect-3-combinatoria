package combinatoria

import (
	"fmt"
	"math"

	"github.com/gookit/color"
)

//FUNCIÓN PARA CALCULAR FACTORIAL
//RECIBE UN SOLO NUMERO ENTERO Y RETORNA UN ENTERO
func Factorial(num int) int {
	var result int = 1
	if num < 0 {
		fmt.Println("Numero debe ser mayor que cero, intenta nuevamente")
	} else if num == 1 {
		return num
	} else {
		for i := 1; i <= num; i++ {
			result *= i
		}
	}
	return result
}

//FUNCION PARA PERMUTACIÓN SIN REPETICIÓN
//RECIBE UN SOLO NUMERO ENTERO Y RETORNA UN ENTERO
func PermutacionSinRepe(num int) int {
	result := Factorial(num)
	return result
}

//FUNCION PARA PERMUTACION CON REPETICION
//ES UNA VARADIC FUNCION QUE RECIBE CUANTOS VALORES EL USUARIO QUIERA INGRESAR, RETORNA UN FLOAT64
func PermutacionConRepe(cantidad ...int) float64 {
	result := 1.0
	var n int
	color.Print("Ingresa <suc>n</>: ")
	fmt.Scanln(&n)
	//var n int = 6
	multFact := 1
	for _, num := range cantidad {
		multFact = multFact * Factorial(num)
	}
	result = float64(Factorial(n) / multFact)
	return result
}

//FUNCIÓN PARA COMBINACIÓN SIN REPETECIÓN
//RECIBE n DE TIPO ENTERO Y r DE TIPO ENTERO, RETORNA UN ENTERO
func CombinacionSinRepe(n, r int) int {
	result := Factorial(n) / (Factorial(n-r) * Factorial(r))
	return result
}

//FUNCIÓN PARA COMBINACIÓN CON REPETICIÓN
//RECIBE n DE TIPO ENTERO Y r DE TIPO ENTERO, RETORNA UN ENTERO
func CombinacionConRepe(n, r int) int {
	result := Factorial(n+r-1) / (Factorial(r) * Factorial(n-1))
	return result
}

//FUNCIÓN PARA VARIACIONES SIN REPETICIÓN
//RECIBE n DE TIPO ENTERO Y r DE TIPO ENTERO, RETORNA UN ENTERO
func VariacionesSinRepe(n, r int) int {
	result := Factorial(n) / Factorial(n-r)
	return result
}

//FUNCIÓN PARA VARIACIONES CON REPETICIÓN
//RECIBE n DE TIPO ENTERO Y r DE TIPO ENTERO, RETORNA UN FLOAT64
func VariacionesConRepe(n, r float64) float64 {
	result := math.Pow(n, r)
	return result
}

//FUNCIÓN PARA EMPAREJAMIENTOS
func Emparejamiento(n, m int) int {
	result := 0
	nxm := n * m
	nxmFact := Factorial(nxm)
	mf := Factorial(m)
	nf := Factorial(n)
	result = int(nxmFact) / (int(math.Pow(float64(mf), float64(n)) * float64(nf)))
	return result
}
