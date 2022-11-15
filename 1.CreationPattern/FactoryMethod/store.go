package factoryMethod

import "io"

type Store interface {
	Open(path string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case DiskStorage:
	case TempStorage:
	case MemoryStorage:
	}
	return nil
}
