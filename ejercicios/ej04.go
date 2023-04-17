package ejercicios

// Escriba un método recursivo que tome una cadena y
// devuelva como resultado la cadena invertida.
func Invertir(cadena string) string {
	if len(cadena) <= 1 {
		return cadena
	}
	invertido := string(cadena[len(cadena)-1])

	return invertido + Invertir(cadena[:len(cadena)-1])
}
