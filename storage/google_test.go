package storage_test

import (
	"context"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/tj/assert"

	"github.com/matsilva/goinstall"
	"github.com/matsilva/goinstall/storage"
)

// newStorage helper.
func newStorage(t testing.TB) goinstall.Storage {
	skipWithoutGoogleCredentials(t)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		t.Fatalf("error creating storage client: %s", err)
	}

	return storage.NewGoogle(&storage.GoogleOptions{
		Client: client,
		Bucket: "goinstall",
	})
}

// Test object creation.
func TestGoogle_Create(t *testing.T) {
	s := newStorage(t)
	ctx := context.Background()

	bin := goinstall.Binary{
		Path:    "github.com/tj/node-prune",
		Version: "v1.0.0",
		OS:      "darwin",
		Arch:    "amd64",
	}

	err := s.Create(ctx, strings.NewReader("Hello World"), bin)
	assert.NoError(t, err)
}

// Test fetching an object.
func TestGoogle_Get(t *testing.T) {
	s := newStorage(t)
	ctx := context.Background()

	t.Run("exists", func(t *testing.T) {
		bin := goinstall.Binary{
			Path:    "github.com/tj/node-prune",
			Version: "v1.0.0",
			OS:      "darwin",
			Arch:    "amd64",
		}

		r, err := s.Get(ctx, bin)
		assert.NoError(t, err)
		defer r.Close()

		b, err := ioutil.ReadAll(r)
		assert.NoError(t, err)
		assert.Equal(t, "Hello World", string(b))
	})

	t.Run("missing", func(t *testing.T) {
		bin := goinstall.Binary{
			Path:    "github.com/tj/node-prune",
			Version: "v2.1.0",
			OS:      "darwin",
			Arch:    "amd64",
		}

		_, err := s.Get(ctx, bin)
		assert.Equal(t, goinstall.ErrObjectNotFound, err)
	})
}

// skipWithoutGoogleCredentials skips the tests unless GCP credentials are present.
func skipWithoutGoogleCredentials(t testing.TB) {
	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		t.SkipNow()
	}
}
