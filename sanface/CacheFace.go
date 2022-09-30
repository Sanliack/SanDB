package sanface

type CacheFace interface {
	GetLen() int
	Put(dbname, k string, v []byte)
	Get(database, key string) ([]byte, bool)
	Del(name, key string)
	Clean(name string)
	TestDuBug()
	//FindData()
}
