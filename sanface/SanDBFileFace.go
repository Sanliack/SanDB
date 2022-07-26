package sanface

type SanDBFileFace interface {
	Read(int64) (EntryFace, error)
	GetOffset() int64
	Write(EntryFace) error
}
