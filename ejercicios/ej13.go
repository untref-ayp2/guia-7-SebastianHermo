package ejercicios

// Escriba un método recursivo para determinar si
// un elemento está en un arreglo utilizando búsqueda binaria.
func BusquedaBinaria(arreglo []int, elemento int) bool { //Complejidad O(n)
	if len(arreglo) < 1 {
		return false
	}
	if arreglo[0] == elemento {
		return true
	}
	return BusquedaBinaria(arreglo[1:], elemento)
}
