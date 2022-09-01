package sanface

type WorkerTranDataFace interface {
	GetData() TranDataFace
	GetConn() ConnFace
}
