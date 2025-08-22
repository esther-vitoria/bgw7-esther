package application

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplicationDefault_NewApplicationDefault(t *testing.T) {
	t.Run("Create a new ApplicationDefault without parameters", func(t *testing.T) {
		application := NewApplicationDefault(nil)
		require.NotNil(t, application)
	})

	t.Run("Create a new ApplicationDefault with parameters", func(t *testing.T) {
		application := NewApplicationDefault(&ConfigApplicationDefault{
			Addr: ":8080",
		})
		require.NotNil(t, application)
	})
}

func TestApplicationDefault_TearDown(t *testing.T) {
	t.Run("Tear down the application", func(t *testing.T) {
		application := NewApplicationDefault(nil)
		err := application.TearDown()
		require.NoError(t, err)
	})
}

func TestApplicationDefault_SetUp(t *testing.T) {
	// Criar arquivo temporário para teste
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "products.json")
	os.WriteFile(tmpFile, []byte("[]"), 0644) // exemplo de JSON vazio

	cfg := &ConfigApplicationDefault{
		ProductJSONPath: tmpFile,
	}

	application := NewApplicationDefault(cfg)
	err := application.SetUp(cfg)
	require.NoError(t, err)
}
