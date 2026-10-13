package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestCodigoSaida(t *testing.T) {
	casos := []struct {
		nome     string
		err      error
		esperado int
	}{
		{"sem erro", nil, 0},
		{"erro generico", errors.New("falha qualquer"), 1},
		{"erro de uso", ErrUso, 2},
		{"erro de validacao", &ErroValidacao{}, 3},
		{"uso embrulhado", fmt.Errorf("contexto: %w", ErrUso), 2},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			if got := codigoSaida(c.err); got != c.esperado {
				t.Errorf("codigoSaida(%v) = %d; esperado %d", c.err, got, c.esperado)
			}
		})
	}
}

func TestCarregarInexistente(t *testing.T) {
	arq := filepath.Join(t.TempDir(), "inexistente.json")
	got, err := carregar(arq)
	if err != nil {
		t.Fatalf("carregar retornou erro inesperado: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("esperado slice vazio, obteve len=%d", len(got))
	}
}

func TestCarregarJSON(t *testing.T) {
	arq := filepath.Join(t.TempDir(), "tarefas.json")
	dados := []byte(`[{"titulo":"comprar café","peso":1},{"titulo":"estudar Go","peso":2}]`)
	if err := os.WriteFile(arq, dados, 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := carregar(arq)
	if err != nil {
		t.Fatalf("carregar falhou inesperadamente: %v", err)
	}

	esperado := []Tarefa{
		{Titulo: "comprar café", Peso: 1},
		{Titulo: "estudar Go", Peso: 2},
	}
	if !slices.Equal(got, esperado) {
		t.Errorf("got %v, esperado %v", got, esperado)
	}
}

func TestCarregarJSONInvalido(t *testing.T) {
	arq := filepath.Join(t.TempDir(), "corrompido.json")
	dados := []byte("isto nao e um json valido")
	if err := os.WriteFile(arq, dados, 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := carregar(arq)
	if err == nil {
		t.Fatalf("esperado erro ao decodificar JSON inválido, mas obteve nil (got=%v)", got)
	}
}
