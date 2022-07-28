package sanface

type ClientControlFace interface {
	Put([]byte, []byte) error
	Get([]byte) ([]byte, error)
	Del([]byte) error
	Clear() error
}
