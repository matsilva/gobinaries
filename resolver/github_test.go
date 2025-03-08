package resolver_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-github/v28/github"
	"github.com/tj/assert"
	"golang.org/x/oauth2"

	"github.com/matsilva/goinstall"
	"github.com/matsilva/goinstall/resolver"
)

// newResolver returns a new GitHub resolver.
func newResolver(t testing.TB) goinstall.Resolver {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		t.Skip("GITHUB_TOKEN environment variable is required")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)

	return &resolver.GitHub{
		Client: github.NewClient(oauth2.NewClient(ctx, ts)),
	}
}

// Test resolver.
func TestGitHub_Resolve(t *testing.T) {
	r := newResolver(t)

	t.Run("exact match", func(t *testing.T) {
		v, err := r.Resolve("tj", "d3-bar", "v1.8.0")
		assert.NoError(t, err)
		assert.Equal(t, "v1.8.0", v)
	})

	t.Run("exact match without leading v", func(t *testing.T) {
		v, err := r.Resolve("tj", "d3-bar", "1.8.0")
		assert.NoError(t, err)
		assert.Equal(t, "v1.8.0", v)
	})

	t.Run("major wildcard match", func(t *testing.T) {
		v, err := r.Resolve("tj", "d3-bar", "1.x")
		assert.NoError(t, err)
		assert.Equal(t, "v1.8.0", v)
	})

	t.Run("minor wildcard match", func(t *testing.T) {
		v, err := r.Resolve("tj", "d3-bar", "1.6.x")
		assert.NoError(t, err)
		assert.Equal(t, "v1.6.0", v)
	})

	t.Run("minor match", func(t *testing.T) {
		v, err := r.Resolve("tj", "d3-bar", "1.6")
		assert.NoError(t, err)
		assert.Equal(t, "v1.6.0", v)
	})

	t.Run("master", func(t *testing.T) {
		v, err := r.Resolve("tj", "d3-bar", "master")
		assert.NoError(t, err)
		assert.Equal(t, "v1.8.0", v)
	})
}
