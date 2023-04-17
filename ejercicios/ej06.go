package ejercicios

// Escriba un método recursivo que calcule el máximo
// común divisor entre dos números enteros.
// Nota: Se puede usar el algoritmo de Euclides para
// resolver este problema.
func MCD(a, b int) int {
	mayor, menor := a, b
	if a < b {
		mayor, menor = b, a
	}
	if mayor%menor == 0 {
		return menor
	}
	menor, mayor = mayor%menor, menor
	return MCD(mayor, menor)
}
