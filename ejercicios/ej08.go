package ejercicios

// Escribir un método recursivo que devuelva el
// cociente y el resto de la división entera mediante
// restas sucesivas
func DivisionEntera(dividendo, divisor int) (cociente, resto int) {
	if dividendo == 0 || divisor == 0 {
		return 0, 0
	}
	if dividendo%divisor == 0 {
		return dividir(dividendo, divisor), 0
	} else {
		return dividir(dividendo, divisor), 1
	}
}

func dividir(dividendo, divisor int) int {
	if dividendo-divisor < 0 {
		return 0
	} else {
		return 1 + dividir(dividendo-divisor, divisor)
	}
}
