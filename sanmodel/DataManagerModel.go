package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

type DataManagerModel struct {
	Name      string
	SanDBFile sanface.FileFace
	IndexMap  map[string]int64
	DMLock    sync.RWMutex
}

func (d *DataManagerModel) Put(key []byte, val []byte) error {
	d.DMLock.Lock()
	defer d.DMLock.Unlock()

	entry := NewEntryModel(key, val, Str_Put)
	err := d.SanDBFile.Write(entry)
	if err != nil {
		fmt.Println("[Error] Conn User SanDBFile Func <Write> appear Error", err)
		return err
	}
	d.IndexMap[string(key)] = d.SanDBFile.GetOffset() - entry.GetSize()
	return nil
}

func (d *DataManagerModel) Get(key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("len of key is 0")
	}
	d.DMLock.RLock()
	defer d.DMLock.RUnlock()
	//c.ConnLock.RLock()
	//defer c.ConnLock.RUnlock()
	offset, ok := d.IndexMap[string(key)]
	if !ok {
		fmt.Println("[Info] key no exist")
		return nil, errors.New("key no exist")
	}
	entry, err := d.SanDBFile.Read(offset)
	if err != nil && err != io.EOF {
		//fmt.Println("============================")
		return nil, err
	} else if err == io.EOF {
		return nil, nil
	}
	return entry.GetVal(), nil
}

func (d *DataManagerModel) Del(key []byte) error {
	if len(key) == 0 {
		return nil
	}
	if _, ok := d.IndexMap[string(key)]; !ok {
		return nil
	}
	d.DMLock.Lock()
	defer d.DMLock.Unlock()
	delentry := NewEntryModel(key, nil, Str_Del)
	err := d.SanDBFile.Write(delentry)
	if err != nil {
		fmt.Println("[Warning] ConnModel use Func <Del> Appear error:", err)
		return err
	}
	delete(d.IndexMap, string(key))
	return nil
}

func (d *DataManagerModel) Clean() error {
	d.DMLock.Lock()
	defer d.DMLock.Unlock()
	d.IndexMap = make(map[string]int64)
	err := d.SanDBFile.Clean()
	if err != nil {
		fmt.Println("[Error] ConnModel Try To Clean fileData appear Error:", err)
		return err
	}
	fmt.Println("clean over ? map:", d.IndexMap)
	return nil
}

func (d *DataManagerModel) MergeFile() error {
	if d.SanDBFile.GetOffset() == 0 {
		return nil
	}
	d.DMLock.Lock()
	defer d.DMLock.Unlock()

	var offset int64
	var newEntrys []sanface.EntryFace
	for {
		entry, err := d.SanDBFile.Read(offset)
		if err != nil {
			if err == io.EOF {
				fmt.Println("[info] ConnModel Read EOF exit ")
				break
			} else {
				fmt.Println("[Error] ConnModel MergeFile user Func <SanDBFile.Read> appear Error", err)
				return err
			}
		}
		off, ok := d.IndexMap[string(entry.GetKey())]
		if !(!ok || entry.GetMask() == Str_Del || off != offset) {
			fmt.Println("append:=", string(entry.GetKey())+":"+string(entry.GetVal()), "msgtype :", entry.GetMask(), "off:", off, "offset:", offset)
			newEntrys = append(newEntrys, entry)
		}
		offset += entry.GetSize()
	}

	var newfile sanface.FileFace
	var err error

	if len(newEntrys) > 0 {
		newfile, err = NewSanDBFileModel(conf.ConfigObj.SanDBFileMergePath + d.Name + ".merge.data")
		if err != nil {
			fmt.Println("[Error] ConnModel-MergeFile user Func <NewSanDBFileModel> appear error:", err)
			return err
		}
		for _, entry := range newEntrys {
			newoffset := newfile.GetOffset()
			err = newfile.Write(entry)
			if err != nil {

			}
			d.IndexMap[string(entry.GetKey())] = newoffset
		}
	}
	olddatafilename := d.SanDBFile.GetFile().Name()
	err = d.SanDBFile.GetFile().Close()
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

	d.SanDBFile = newfile
	return nil
}

func (d *DataManagerModel) InitMap() {
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
		if mask == Str_Del {
			delete(d.IndexMap, string(entry.GetKey()))
			offset += entry.GetSize()
			continue
		}
		key := entry.GetKey()
		d.IndexMap[string(key)] = offset
		offset += entry.GetSize()
	}
	fmt.Printf("DataManagerModel %s init Map sucess len of IndexMap:%d IndexMap KV:%v \n", d.Name, len(d.IndexMap), d.IndexMap)
}

func NewDataManagerModel(name string) (sanface.DataManagerFace, error) {
	fmt.Println("name:=====", conf.ConfigObj.SanDBFilePath+name+".str.data")
	filemodel, err := NewSanDBFileModel(conf.ConfigObj.SanDBFilePath + name + ".str.data")
	if err != nil {
		fmt.Println("")
		return nil, err
	}
	dm := &DataManagerModel{
		Name:      name,
		SanDBFile: filemodel,
		IndexMap:  make(map[string]int64),
	}
	dm.InitMap()
	return dm, nil
}
