package sanface

type WorkerFace interface {
	Start()
	Stop()
	ReStart()
	AddMsg(data WorkerTranDataFace)
}
