package sanface

type ConnFace interface {
	Start()
	Listen()
	Stop()
}
