package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"fmt"
	"io"
	"os"
	"sync"
)

type SetManagerModel struct {
	Name      string
	SanDBFile sanface.FileFace
	IndexMap  map[string]map[string]int64
	DMLock    sync.RWMutex
}

func (s *SetManagerModel) Sadd(key []byte, val []byte) error {
	s.DMLock.Lock()
	defer s.DMLock.Unlock()
	if s.IndexMap[string(key)] != nil {
		_, ok := s.IndexMap[string(key)][string(val)]
		if ok {
			return nil
		}
	}
	entry := NewEntryModel(key, val, Set_Add)
	err := s.SanDBFile.Write(entry)
	if err != nil {
		fmt.Println("[Error]SetManagerModel try to write file appear error,", err)
		return err
	}
	if s.IndexMap[string(key)] == nil {
		s.IndexMap[string(key)] = make(map[string]int64, 0)
	}
	s.IndexMap[string(key)][string(val)] = s.SanDBFile.GetOffset() - entry.GetSize()
	return nil
}

func (s *SetManagerModel) Scard(key []byte) int {
	_, ok := s.IndexMap[string(key)]
	if !ok {
		return 0
	}
	return len(s.IndexMap[string(key)])
}

func (s *SetManagerModel) Smember(key []byte) ([][]byte, error) {
	s.DMLock.RLock()
	defer s.DMLock.RUnlock()
	//keynum := s.Scard(key)
	_, ok := s.IndexMap[string(key)]
	if !ok {
		fmt.Println("[info] Client Search no exist key")
		return nil, nil
	}
	sets := make([][]byte, len(s.IndexMap[string(key)]))
	no := 0
	for _, aset := range s.IndexMap[string(key)] {
		set, err := s.SanDBFile.Read(aset)
		if err != nil {
			fmt.Println("[Warning] Smenber Read Error ", err)
			return nil, err
		}
		sets[no], _ = set.Encode()
		no++
	}
	return sets, nil
}

func (s *SetManagerModel) Spop(key []byte, val []byte) error {
	s.DMLock.Lock()
	defer s.DMLock.Unlock()
	keynum := s.Scard(key)
	if keynum == 0 {
		fmt.Println("[info] Client Search no exist key")
		return nil
	}
	_, flag := s.IndexMap[string(key)][string(val)]
	if !flag {
		return nil
	}

	delentry := NewEntryModel(key, val, Set_Del)
	err := s.SanDBFile.Write(delentry)
	if err != nil {
		fmt.Println("")
	}
	delete(s.IndexMap[string(key)], string(val))
	return nil
}

func (s *SetManagerModel) DelByKey(key []byte) error {
	s.DMLock.Lock()
	defer s.DMLock.Unlock()
	_, flag := s.IndexMap[string(key)]
	if !flag {
		return nil
	} else {
		for val, _ := range s.IndexMap[string(key)] {
			delentry := NewEntryModel(key, []byte(val), Set_Del)
			err := s.SanDBFile.Write(delentry)
			if err != nil {
				fmt.Println("")
				return err
			}
		}
		delete(s.IndexMap, string(key))
	}
	return nil
}

func (s *SetManagerModel) MergeFile() error {
	if s.SanDBFile.GetOffset() == 0 {
		return nil
	}
	s.DMLock.Lock()
	defer s.DMLock.Unlock()

	var newEntrys []sanface.EntryFace
	for _, vmap := range s.IndexMap {
		if len(vmap) == 0 {
			continue
		}
		for _, offset := range vmap {
			entry, err := s.SanDBFile.Read(offset)
			if err != nil {
				fmt.Println("mergeFile search error")
				return err
			}
			newEntrys = append(newEntrys, entry)
		}
	}

	var newfile sanface.FileFace
	var err error
	var newmap = map[string][]int64{}

	if len(newEntrys) > 0 {
		newfile, err = NewSanDBFileModel(conf.ConfigObj.SanDBFileMergePath + s.Name + ".set.merge.data")
		if err != nil {
			fmt.Println("[Error] ConnModel-MergeFile user Func <NewSanDBFileModel> appear error:", err)
			return err
		}
		for _, entry := range newEntrys {
			newoffset := newfile.GetOffset()
			err = newfile.Write(entry)
			if err != nil {

			}
			if newmap[string(entry.GetKey())] == nil {
				newmap[string(entry.GetKey())] = make([]int64, 1)
				newmap[string(entry.GetKey())][0] = newoffset
			} else {
				newmap[string(entry.GetKey())] = append(newmap[string(entry.GetKey())], newoffset)
			}
		}
	}
	olddatafilename := s.SanDBFile.GetFile().Name()
	err = s.SanDBFile.GetFile().Close()
	if err != nil {
		fmt.Println("[Error] ConnModel Close OldDataFile appear Error:", err)
		return err
	}

	newfileneme := newfile.GetFile().Name()
	err = newfile.GetFile().Close()
	if err != nil {
		fmt.Println("[Error] ConnModel Close NewDataFile appear Error:", err)
		return err
	}
	err = os.Remove(olddatafilename)
	if err != nil {
		fmt.Println("[Error] ConnModel Merge NewDataFile Remove old file appear Error:", err)
		return err
	}

	err = os.Rename(newfileneme, olddatafilename)
	if err != nil {
		fmt.Println("[Error] ConnModel Merge NewDataFile Change Name instead OldDataFile appear Error:", err)
		return err
	}
	s.SanDBFile = newfile
	return nil
}

func (s *SetManagerModel) Clean() error {
	s.DMLock.Lock()
	err := s.SanDBFile.Clean()
	if err != nil {
		fmt.Println("[Error] SetManagerModel Clean error", err)
		return err
	}
	s.IndexMap = make(map[string]map[string]int64)
	s.DMLock.Unlock()
	if err != nil {
		fmt.Println("clean error")
		return err
	}
	return nil
}

func (s *SetManagerModel) SIsMember(key []byte, val []byte) bool {
	_, ok := s.IndexMap[string(key)][string(val)]
	if ok {
		return true
	}
	return false
}

func (d *SetManagerModel) SetInitMap() {
	var offset int64
	for {
		entry, err := d.SanDBFile.Read(offset)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("[Warning] Conn init IndexMap appear error", err)
			break
		}
		mask := entry.GetMask()
		if mask == Set_Del {
			if _, ok := d.IndexMap[string(entry.GetKey())]; ok {
				delete(d.IndexMap[string(entry.GetKey())], string(entry.GetVal()))
			}
			offset += entry.GetSize()
			continue
		}
		key := entry.GetKey()
		if d.IndexMap[string(key)] == nil {
			d.IndexMap[string(key)] = make(map[string]int64)
		}
		d.IndexMap[string(key)][string(entry.GetVal())] = offset
		offset += entry.GetSize()
	}
	fmt.Printf("DataManagerModel %s init Map sucess len of IndexMap:%d IndexMap KV:%v \n", d.Name, len(d.IndexMap), d.IndexMap)
}

func NewSetManagerModel(name string) (sanface.SetManagerFace, error) {
	filemodel, err := NewSanDBFileModel(conf.ConfigObj.SanDBFilePath + name + ".sets.data")
	if err != nil {
		fmt.Println("")
		return nil, err
	}
	dm := &SetManagerModel{
		Name:      name,
		SanDBFile: filemodel,
		IndexMap:  make(map[string]map[string]int64),
	}
	dm.SetInitMap()
	return dm, nil
}
