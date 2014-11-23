package backend

import "errors"

type Type string

const (
	Data     Type = "data"
	Key           = "key"
	Lock          = "lock"
	Snapshot      = "snapshot"
	Tree          = "tree"
	Map           = "map"
)

const (
	BackendVersion = 1
)

var (
	ErrAlreadyPresent = errors.New("blob is already present in backend")
)

type Server interface {
	Create(Type, []byte) (ID, error)
	Get(Type, ID) ([]byte, error)
	List(Type) (IDs, error)
	Test(Type, ID) (bool, error)
	Remove(Type, ID) error
	Version() uint

	Close() error

	Location() string
}