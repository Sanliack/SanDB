package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"encoding/binary"
)

type EntryModel struct {
	Key     []byte
	Val     []byte
	keySize uint32
	ValSize uint32
	Mark    uint16
}

const (
	Put uint16 = iota
	Del
	Get
	Msg
)

func (e *EntryModel) GetSize() int64 {
	return int64(conf.ConfigObj.EntryHeaderSize + int(e.keySize+e.ValSize))
}

func (e *EntryModel) GetKey() []byte {
	return e.Key
}

func (e *EntryModel) GetVal() []byte {
	return e.Val
}

func (e *EntryModel) GetMask() uint16 {
	return e.Mark
}

func (e *EntryModel) GetKeySize() uint32 {
	return e.keySize
}

func (e *EntryModel) SetKey(key []byte) {
	e.Key = key
}

func (e *EntryModel) SetVal(val []byte) {
	e.Val = val
}

func (e *EntryModel) GetValSize() uint32 {
	return e.ValSize
}

func (e *EntryModel) SetMark(mark uint16) {
	e.Mark = mark
}

func (e *EntryModel) Encode() ([]byte, error) {
	buf := make([]byte, e.GetSize())
	binary.BigEndian.PutUint32(buf[0:4], e.keySize)
	binary.BigEndian.PutUint32(buf[4:8], e.ValSize)
	binary.BigEndian.PutUint16(buf[8:10], e.Mark)
	copy(buf[conf.ConfigObj.EntryHeaderSize:conf.ConfigObj.EntryHeaderSize+int(e.keySize)], e.Key)
	copy(buf[conf.ConfigObj.EntryHeaderSize+int(e.keySize):], e.Val)
	return buf, nil
}

func DecodeEntry(buf []byte) (sanface.EntryFace, error) {
	newentry := &EntryModel{}
	newentry.keySize = binary.BigEndian.Uint32(buf[:4])
	newentry.ValSize = binary.BigEndian.Uint32(buf[4:8])
	newentry.Mark = binary.BigEndian.Uint16(buf[6:10])
	return newentry, nil
}

func NewEntryModel(key, val []byte, mark uint16) sanface.EntryFace {
	return &EntryModel{
		Key:     key,
		Val:     val,
		keySize: uint32(len(key)),
		ValSize: uint32(len(val)),
		Mark:    mark,
	}
}
