package sanface

type ConnFace interface {
	Start()
	GetIndexMap() map[string]int64
	GetSanDBFile() SanDBFileFace
	Put([]byte, []byte) error
	Get([]byte) ([]byte, error)
	Del([]byte) error
	MergeFile() error
}
