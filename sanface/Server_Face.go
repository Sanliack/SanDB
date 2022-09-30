package sanface

type Server interface {
	Start()
	Server()
	Stop()
	GetVersion() string
	GetConnNums() int
	GetCacheManager() CacheFace
	GetDataManager(string) (DataManagerFace, error)
	GetSetManager(string) (SetManagerFace, error)
	AddMsgToMsgQueue(data WorkerTranDataFace)
}
