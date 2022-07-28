package sanface

type ClientFace interface {
	Start()
	Stop()
	Server()
	Connect(tcpADdr string) (ClientControlFace, error)
}
