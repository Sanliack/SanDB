package sanface

type EntryFace interface {
	GetSize() int64
	Encode() ([]byte, error)
	GetKeySize() uint32
	GetValSize() uint32
	GetKey() []byte
	GetVal() []byte
	GetMask() uint16
	SetKey(key []byte)
	SetVal(key []byte)
	SetMark(uint16)
}
