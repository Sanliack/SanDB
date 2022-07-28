package sanface

type ConnFace interface {
	Start()
	GetIndexMap() map[string]int64
	GetSanDBFile() FileFace
	Put([]byte, []byte) error
	Get([]byte) ([]byte, error)
	Del([]byte) error
	MergeFile() error
	Listen()
	Stop()
}
