package sanface

type WorkerFace interface {
	Start()
	Stop()
	AddMsg(data WorkerTranDataFace)
}
