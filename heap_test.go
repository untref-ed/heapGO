package heap

import (
	"testing"
)

func TestHeapMin(t *testing.T) {
	// Creamos un Heap Min de enteros
	h := NuevoHeapMin[int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	})
	// Insertamos algunos elementos
	h.Insertar(5)
	h.Insertar(3)
	h.Insertar(7)
	h.Insertar(1)
	h.Insertar(2)
	h.Insertar(4)
	h.Insertar(6)
	// Extraemos los elementos
	for i := 1; i <= 7; i++ {
		e, err := h.Extraer()
		if err != nil {
			t.Error(err)
		}
		if e != i {
			t.Errorf("Se esperaba %d y se obtuvo %d", i, e)
		}
	}
	// Extraer de un Heap vacío
	_, err := h.Extraer()
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

func TestHeapMax(t *testing.T) {
	// Creamos un Heap Max de enteros
	h := NuevoHeapMax[int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	})
	// Insertamos algunos elementos
	h.Insertar(5)
	h.Insertar(3)
	h.Insertar(7)
	h.Insertar(1)
	h.Insertar(2)
	h.Insertar(4)
	h.Insertar(6)
	// Extraemos los elementos
	for i := 7; i >= 1; i-- {
		e, err := h.Extraer()
		if err != nil {
			t.Error(err)
		}
		if e != i {
			t.Errorf("Se esperaba %d y se obtuvo %d", i, e)
		}
	}
	// Extraer de un Heap vacío
	_, err := h.Extraer()
	if err == nil {
		t.Error("Se esperaba un error")
	}
}
