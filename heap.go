package heap

import (
	"fmt"
)

// Heap es una estructura de datos que permite almacenar elementos
// y recuperarlos en orden de prioridad.
// Soporta Heap Min y Heap Max.
type Heap[T any] struct {
	// Arreglo de elementos
	heap []T
	// Función de comparación
	// Para comparar dos elementos del heap
	comp func(T, T) int
	// Es Heap Min?
	min int
}

// Constructor interno
func nuevoHeap[T any](comp func(T, T) int, min int) *Heap[T] {
	h := new(Heap[T])

	h.comp = comp
	h.min = min

	return h
}

// NuevoHeapMin crea un Heap de mínimos.
// La función de comparación debe devolver un valor negativo si a < b,
// cero si a == b, y un valor positivo si a > b.
func NuevoHeapMin[T any](comp func(T, T) int) *Heap[T] {
	return nuevoHeap[T](comp, -1)
}

// NuevoHeapMax crea un Heap de máximos.
// La función de comparación debe devolver un valor negativo si a < b,
// cero si a == b, y un valor positivo si a > b.
func NuevoHeapMax[T any](comp func(T, T) int) *Heap[T] {
	return nuevoHeap[T](comp, 1)
}

// Insertar agrega un elemento al Heap.
func (h *Heap[T]) Insertar(e T) {
	h.heap = append(h.heap, e)
	h.upHeap(len(h.heap) - 1)
}

// Extraer devuelve el elemento de mayor prioridad y lo elimina del Heap.
// Si el Heap está vacío, devuelve un error.
func (h *Heap[T]) Extraer() (T, error) {
	if h.EstaVacio() {
		var nulo T
		return nulo, fmt.Errorf("Heap vacío")
	}

	e := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	h.downHeap(0)

	return e, nil
}

// upHeap sube un elemento en el Heap hasta que se cumpla la propiedad de orden.
func (h *Heap[T]) upHeap(i int) {
	if i == 0 {
		return
	}

	padre := (i - 1) / 2

	if h.comp(h.heap[i], h.heap[padre])*h.min >= 1 {
		h.heap[i], h.heap[padre] = h.heap[padre], h.heap[i]
		h.upHeap(padre)
	}
}

// downHeap baja un elemento en el Heap hasta que se cumpla la propiedad de orden.
func (h *Heap[T]) downHeap(i int) {
	hi := 2*i + 1

	if hi >= len(h.heap) {
		return
	}

	if hi+1 < len(h.heap) && h.comp(h.heap[hi+1], h.heap[hi])*h.min >= 1 {
		hi++
	}

	if h.comp(h.heap[hi], h.heap[i])*h.min >= 1 {
		h.heap[i], h.heap[hi] = h.heap[hi], h.heap[i]
		h.downHeap(hi)
	}
}

// EstaVacio devuelve true si el Heap está vacío.
func (h *Heap[T]) EstaVacio() bool {
	return len(h.heap) == 0
}

// Devuelve una cadena con la representación del Heap
func (h *Heap[T]) String() string {
	return fmt.Sprintf("%v", h.heap)
}
