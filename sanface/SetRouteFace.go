package sanface

type SetRouteFace interface {
	Sadd(trandata TranDataFace) error
	Scard(trandata TranDataFace) error
	Smember(trandata TranDataFace) error
	SIsmember(trandata TranDataFace) error
	Spop(trandata TranDataFace) error
	MergeFile(trandata TranDataFace) error
	DelByKey(trandata TranDataFace) error
	Clean() error
}
