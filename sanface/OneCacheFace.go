package sanface

type OneCacheFace interface {
	Put(key string, val []byte)
	Get(key string) ([]byte, bool)
	TestDebug()
}
