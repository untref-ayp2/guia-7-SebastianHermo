package ejercicios

// Escribir un método recursivo que reciba 2 enteros
// n y b y devuelva true si n es potencia de b.
// Por ejemplo: esPotencia(8, 2) devuelve true.
func EsPotencia(n, b int) bool {
	if n == b || n == 1 {
		return true
	}
	if b > n {
		return false
	}
	return EsPotencia(n, b+b)
}
