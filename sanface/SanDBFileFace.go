package sanface

import "os"

type SanDBFileFace interface {
	Read(int64) (EntryFace, error)
	GetOffset() int64
	GetFile() *os.File
	Write(EntryFace) error
}
