package sanface

type SetControlFace interface {
	Sadd(key []byte, val []byte) error
	Scard(key []byte) (int, error)
	Smember(key []byte) ([][]byte, error)
	Spop(key []byte, val []byte) error
	SIsMember(key []byte, val []byte) (bool, error)
	DelByKey(key []byte) error
	MergeFile() error
	Clean() error
}
