package ejercicios

// Escribir un método recursivo que dado un número
// entero decimal devuelva su equivalente en
// binario en forma de string
func DecimalABinario(n int) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}
	if n%2 == 1 {
		return DecimalABinario(n/2) + "1"
	}
	if n%2 == 0 {
		return DecimalABinario(n/2) + "0"
	}
	return ""
}
