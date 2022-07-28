package sanface

import "os"

type FileFace interface {
	Read(int64) (EntryFace, error)
	GetOffset() int64
	GetFile() *os.File
	Write(EntryFace) error
}
