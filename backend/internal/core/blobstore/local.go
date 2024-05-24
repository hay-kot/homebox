package blobstore

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hay-kot/homebox/backend/pkgs/pathlib"
	"github.com/rs/zerolog/log"
)

// localBlobStore is a blob store implementation backed by the local filesystem.
// Blob R/W operations translate to local file create, read, and write operations.
type localBlobStore struct {
	root string
}

// NewLocalBlobStore creates a local blob store rooted at the specified root directory.
// Keys created, written, and deleted are relative to this root directory.
func NewLocalBlobStore(root string) BlobStore {
	err := os.MkdirAll(root, 0o755)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create data directory")
	}

	return &localBlobStore{
		root: root,
	}
}

func (l *localBlobStore) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	return os.Open(l.resolvePath(key))
}

func (l *localBlobStore) Put(ctx context.Context, key string, content io.Reader) (string, error) {
	path := pathlib.Safe(l.resolvePath(key))

	parent := filepath.Dir(path)
	err := os.MkdirAll(parent, 0o755)
	if err != nil {
		return "", err
	}

	f, err := os.Create(path)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(f, content)
	if err != nil {
		return "", err
	}

	return key, nil
}

func (l *localBlobStore) Delete(ctx context.Context, key string) error {
	return os.Remove(l.resolvePath(key))
}

// resolvePath resolves the full path that corresponds to a blob key, taking the root directory
// into account.
func (l *localBlobStore) resolvePath(key string) string {
	// XXX: A previous iteration of the document storage implementation persisted document paths
	// with its fully qualified filesystem path, which included its root directory as a prefix.
	// This compromised relocation resiliency of the attachment storage directory.
	//
	// For example, a root directory of "/usr/share/homebox" and blob key "gid/documents/id"
	// would be persisted with identifier "/usr/share/homebox/gid/documents/id". This would
	// break file integrity if "/usr/share/homebox" were relocated to "/usr/share/homebox2",
	// even if the runtime storage root directory were changed to "/usr/share/homebox2".
	//
	// The current local storage implementation persists blob keys only, independent of the
	// root path, which fixes this capability. However, to preserve backwards compatibility with
	// existing documents written with prior behavior, assume that any blob keys that are
	// prefixed with the root path are already fully resolved/qualified into filesystem paths.
	if strings.HasPrefix(key, l.root) {
		return key
	}

	return filepath.Join(l.root, key)
}
