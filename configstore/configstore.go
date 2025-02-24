// Copyright 2022 Fastly, Inc.

package configstore

import (
	"errors"

	"github.com/fastly/compute-sdk-go/internal/abi/fastly"
)

var (
	// ErrStoreNotFound indicates the named config store doesn't exist.
	ErrStoreNotFound = errors.New("config store not found")

	// ErrStoreNameEmpty indicates the given config store name
	// was empty.
	ErrStoreNameEmpty = errors.New("config store name was empty")

	// ErrStoreNameInvalid indicates the given config store name
	// was invalid.
	ErrStoreNameInvalid = errors.New("config store name contained invalid characters")

	// ErrStoreNameTooLong indicates the given config store name
	// was too long.
	ErrStoreNameTooLong = errors.New("config store name too long")

	// ErrKeyNotFound indicates a key isn't in a config store.
	ErrKeyNotFound = errors.New("key not found")
)

// Store is a read-only representation of a config store.
type Store struct {
	abiDict *fastly.Dictionary
}

// Open returns a config store with the given name. Names are case
// sensitive.
func Open(name string) (*Store, error) {
	d, err := fastly.OpenDictionary(name)
	if err != nil {
		status, ok := fastly.IsFastlyError(err)
		switch {
		case ok && status == fastly.FastlyStatusBadf:
			return nil, ErrStoreNotFound
		case ok && status == fastly.FastlyStatusNone:
			return nil, ErrStoreNameEmpty
		case ok && status == fastly.FastlyStatusUnsupported:
			return nil, ErrStoreNameTooLong
		case ok && status == fastly.FastlyStatusInval:
			return nil, ErrStoreNameInvalid
		default:
			return nil, err
		}
	}
	return &Store{d}, nil
}

// Get returns the item in the config store with the given key.
func (s *Store) Get(key string) (string, error) {
	if s == nil {
		return "", ErrKeyNotFound
	}

	v, err := s.abiDict.Get(key)
	if err != nil {
		status, ok := fastly.IsFastlyError(err)
		switch {
		case ok && status == fastly.FastlyStatusBadf:
			return "", ErrStoreNotFound
		case ok && status == fastly.FastlyStatusNone:
			return "", ErrKeyNotFound
		default:
			return "", err
		}
	}

	return v, nil
}
