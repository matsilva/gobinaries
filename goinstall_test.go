package goinstall

import (
	"testing"
)

func TestBinary(t *testing.T) {
	b := Binary{
		Path:    "github.com/tj/staticgen/cmd/staticgen",
		Module:  "github.com/tj/staticgen",
		Version: "v1.0.0",
		OS:      "linux",
		Arch:    "amd64",
	}

	if b.Path != "github.com/tj/staticgen/cmd/staticgen" {
		t.Errorf("Expected Path to be 'github.com/tj/staticgen/cmd/staticgen', got %q", b.Path)
	}

	if b.Module != "github.com/tj/staticgen" {
		t.Errorf("Expected Module to be 'github.com/tj/staticgen', got %q", b.Module)
	}

	if b.Version != "v1.0.0" {
		t.Errorf("Expected Version to be 'v1.0.0', got %q", b.Version)
	}

	if b.OS != "linux" {
		t.Errorf("Expected OS to be 'linux', got %q", b.OS)
	}

	if b.Arch != "amd64" {
		t.Errorf("Expected Arch to be 'amd64', got %q", b.Arch)
	}
}

func TestErrors(t *testing.T) {
	if ErrObjectNotFound.Error() != "no cloud storage object" {
		t.Errorf("Expected ErrObjectNotFound message to be 'no cloud storage object', got %q", ErrObjectNotFound.Error())
	}

	if ErrNoVersionMatch.Error() != "no matching version" {
		t.Errorf("Expected ErrNoVersionMatch message to be 'no matching version', got %q", ErrNoVersionMatch.Error())
	}

	if ErrNoVersions.Error() != "no versions defined" {
		t.Errorf("Expected ErrNoVersions message to be 'no versions defined', got %q", ErrNoVersions.Error())
	}
}
