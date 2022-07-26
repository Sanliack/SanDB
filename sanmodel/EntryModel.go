package sanmodel

type EntryModel struct {
	Key     []byte
	Val     []byte
	keySize int
	ValSize int
	Mark    uint
}

func NewEntryModel(key, val []byte, mark uint) *EntryModel {
	return &EntryModel{
		Key:     key,
		Val:     val,
		keySize: len(key),
		ValSize: len(val),
		Mark:    mark,
	}
}
