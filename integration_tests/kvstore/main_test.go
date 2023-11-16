//go:build ((tinygo.wasm && wasi) || wasip1) && !nofastlyhostcalls

// Copyright 2023 Fastly, Inc.

package main

import (
	"strings"
	"testing"

	"github.com/fastly/compute-sdk-go/kvstore"
)

func TestKVStore(t *testing.T) {
	store, err := kvstore.Open("example-test-kv-store")
	if err != nil {
		t.Fatal(err)
	}

	hello, err := store.Lookup("hello")
	if err != nil {
		t.Fatal(err)
	}

	if got, want := hello.String(), "world"; got != want {
		t.Errorf("Lookup: got %q, want %q", got, want)
	}

	err = store.Insert("animal", strings.NewReader("cat"))
	if err != nil {
		t.Fatal(err)
	}

	animal, err := store.Lookup("animal")
	if err != nil {
		t.Fatal(err)
	}

	if got, want := animal.String(), "cat"; got != want {
		t.Errorf("Insert: got %q, want %q", got, want)
	}
}
