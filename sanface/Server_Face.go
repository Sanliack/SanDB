package sanface

type Server interface {
	Start()
	Server()
	Stop()
	GetVersion() string
	GetConnNums() int
	GetDataManager(string) (DataManagerFace, error)
	GetSetManager(string) (SetManagerFace, error)
}
