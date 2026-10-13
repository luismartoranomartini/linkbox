package sub

import (
	"os"
	"path/filepath"
	"testing"
)

// func TestPasta(t *testing.T) {
// 	arq := filepath.Join(t.TempDir(), "a.txt")
// 	if err := os.WriteFile(arq, []byte("oi"), 0o644); err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(arq)
// }

func TestPasta(t *testing.T) {
	arq := filepath.Join(t.TempDir(), "a.txt")
	if err := os.WriteFile(arq, []byte("oi"), 0o644); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(arq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(arq, string(b))
}
