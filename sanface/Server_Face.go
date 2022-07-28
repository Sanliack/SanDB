package sanface

type Server interface {
	Start()
	Server()
	Stop()
	GetVersion() string
	GetConnNums() int
}
