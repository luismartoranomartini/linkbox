package sub

import "testing"

func TestExemplo(t *testing.T) {
	t.Run("passa", func(t *testing.T) {})
	t.Run("falha", func(t *testing.T) { t.Fatal("motivo 1") })
	t.Run("erro", func(t *testing.T) { t.Error("motivo 2") })
}
