package sanface

type OneCacheFace interface {
	Put(key string, val []byte)
	Get(key string) ([]byte, bool)
	Del(key string)
	Clean()
	TestDebug()
}
