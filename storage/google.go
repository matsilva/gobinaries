package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"cloud.google.com/go/storage"

	"github.com/matsilva/goinstall"
)

// ErrObjectNotFound is returned from Get when no object is found for the specified key.
var ErrObjectNotFound = errors.New("no cloud storage object")

// GoogleOptions represents the options for Google Cloud Storage.
type GoogleOptions struct {
	// Client is the Google Cloud Storage client.
	Client *storage.Client

	// Bucket is the bucket name.
	Bucket string

	// Prefix is an optional object key prefix.
	Prefix string
}

// Google is a Google Cloud Storage object store for binaries.
type Google struct {
	// Client is the Google Cloud Storage client.
	Client *storage.Client

	// Bucket is the bucket name.
	Bucket string

	// Prefix is an optional object key prefix.
	Prefix string
}

// NewGoogle creates a new Google Cloud Storage client with the given options.
func NewGoogle(opts *GoogleOptions) *Google {
	return &Google{
		Client: opts.Client,
		Bucket: opts.Bucket,
		Prefix: opts.Prefix,
	}
}

// Create an object representing the package's binary.
func (g *Google) Create(ctx context.Context, r io.Reader, bin goinstall.Binary) error {
	key := g.getKey(bin)

	obj := g.Client.Bucket(g.Bucket).Object(key)
	dst := obj.NewWriter(ctx)

	_, err := io.Copy(dst, r)
	if err != nil {
		return fmt.Errorf("copying: %w", err)
	}

	err = dst.Close()
	if err != nil {
		return fmt.Errorf("closing: %w", err)
	}

	return nil
}

// Get returns an object.
func (g *Google) Get(ctx context.Context, bin goinstall.Binary) (io.ReadCloser, error) {
	key := g.getKey(bin)
	obj := g.Client.Bucket(g.Bucket).Object(key)
	r, err := obj.NewReader(ctx)

	if isNotExists(err) {
		return nil, goinstall.ErrObjectNotFound
	}

	return r, nil
}

// getKey returns the object key in the form `<pkg>/<binary>`.
func (g *Google) getKey(bin goinstall.Binary) string {
	dir := g.Prefix + "/" + strings.Replace(bin.Path, "/", "-", -1)
	file := fmt.Sprintf("%s-%s-%s", bin.Version, bin.OS, bin.Arch)
	return dir + "/" + file
}

// isNotExists returns true if the err is present and represents a missing Cloud Storage object.
func isNotExists(err error) bool {
	return err != nil && err.Error() == "storage: object doesn't exist"
}

// NewClient creates a new Google Cloud Storage client.
func NewClient(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx)
}
