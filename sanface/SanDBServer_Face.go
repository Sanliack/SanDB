package sanface

type SanDBServer interface {
	Start()
	Server()
	Stop()
	StringSet()
	StringGet()
	StringGetSet()
}
