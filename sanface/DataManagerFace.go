package sanface

type DataManagerFace interface {
	Put(key []byte, val []byte) error
	Get(key []byte) ([]byte, error)
	Del(key []byte) error
	Clean() error
	MergeFile() error
}
