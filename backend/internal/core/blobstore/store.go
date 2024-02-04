// Package blobstore provides blob storage abstractions for reading and writing binary blobs.
package blobstore

import (
	"io"
)

// BlobStore is an interface that describes a key-value-oriented blob storage backend for arbitrary
// binary objects.
//
// Keys in the blob store are implementation-defined unique identifiers for locating a blob.
type BlobStore interface {
	// Get retrieves a blob by key, returning an io.ReadCloser capable of streaming the blob
	// contents. Callers should close the returned blob to avoid leaks.
	Get(key string) (io.ReadCloser, error)
	// Put creates a new blob with the specified key and contents, and returns a normalized key
	// that can be used for future R/W.
	//
	// Note that the returned key may be identical to that supplied in the original request;
	// the behavior is implementation-defined.
	Put(key string, content io.Reader) (string, error)
	// Delete deletes a blob by key.
	Delete(key string) error
}
