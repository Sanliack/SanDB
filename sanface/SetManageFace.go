package sanface

type SetManagerFace interface {
	Sadd(key []byte, val []byte) error
	Scard(key []byte) int
	Smember(key []byte) ([][]byte, error)
	Spop(key []byte, val []byte) error
	DelByKey(key []byte) error
	MergeFile() error
	Clean() error
	SIsMember(key []byte, val []byte) bool
}
