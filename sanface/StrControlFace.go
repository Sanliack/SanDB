package sanface

type StrControlFace interface {
	Put(key []byte, val []byte) error
	Get(key []byte) ([]byte, error)
	Merge() error
	Del(key []byte) error
	Clean() error
}
