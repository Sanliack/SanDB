package sanface

type CacheFace interface {
	GetLen() int
	Put(dbname, k string, v []byte)
	Get(database, key string) ([]byte, bool)
	TestDuBug()
	//FindData()
}
