package sanface

type TranDataFace interface {
	GetData() []byte
	GetCommId() uint16
	GetLen() uint32
	GetSize() uint32
	Encode() ([]byte, error)
}
