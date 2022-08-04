package sanface

type StrRouteFace interface {
	Del(trandata TranDataFace) error
	Get(trandata TranDataFace) error
	Put(trandata TranDataFace) error
	Clean() error
	Merge() error
}
