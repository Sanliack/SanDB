package sanmodel

import (
	"SanDB/sanface"
	"encoding/binary"
)

type TranDataModel struct {
	TranData  []byte
	CommandId uint16
	Len       uint32
}

func (t *TranDataModel) GetData() []byte {
	return t.TranData
}

func (t *TranDataModel) GetCommId() uint16 {
	return t.CommandId
}

func (t *TranDataModel) GetLen() uint32 {
	return t.Len
}
func (t *TranDataModel) GetSize() uint32 {
	return uint32(len(t.TranData) + 2 + 4)
}

func (t *TranDataModel) Encode() ([]byte, error) {
	buf := make([]byte, t.GetSize())
	binary.BigEndian.PutUint32(buf[0:4], t.Len)
	binary.BigEndian.PutUint16(buf[4:6], t.CommandId)
	copy(buf[6:], t.TranData)
	return buf, nil
}

func DecodeTranData(buf []byte) sanface.TranDataFace {
	tdatalen := binary.BigEndian.Uint32(buf[0:4])
	tdataCommand := binary.BigEndian.Uint16(buf[4:6])
	tdata := buf[6:]
	return &TranDataModel{
		TranData:  tdata,
		CommandId: tdataCommand,
		Len:       tdatalen,
	}
}

func NewTranDataModel(tdata []byte, command uint16) sanface.TranDataFace {
	return &TranDataModel{
		TranData:  tdata,
		CommandId: command,
		Len:       uint32(len(tdata) + 2),
	}
}
