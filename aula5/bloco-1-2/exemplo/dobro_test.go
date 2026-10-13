package exemplo

import "testing"

// func TestDobro(t *testing.T) {
// 	casos := []struct{ entrada, esperado int }{{3, 6}, {0, 0}, {-2, -4}}
// 	for _, c := range casos {
// 		if got := Dobro(c.entrada); got != c.esperado {
// 			t.Errorf("Dobro(%d) = %d, esperado %d", c.entrada, got, c.esperado)
// 		}
// 	}
// }

func TestDobro(t *testing.T) {
	casos := []struct {
		nome              string
		entrada, esperado int
	}{
		{"positivo", 3, 6},
		{"zero", 0, 1},
		{"negativo", -2, -3},
	}
	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			if got := Dobro(c.entrada); got != c.esperado {
				t.Errorf("Dobro(%d) = %d, esperado %d", c.entrada, got, c.esperado)
			}
		})
	}
}
