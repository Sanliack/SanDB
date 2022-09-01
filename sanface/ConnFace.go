package sanface

type ConnFace interface {
	Start()
	Listen()
	Stop()
	SolveTranData(trandata TranDataFace) error
}
