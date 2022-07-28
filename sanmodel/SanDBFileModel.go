package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"fmt"
	"io"
	"os"
)

type SanDBFileModel struct {
	File   *os.File
	Offset int64
}

func (s *SanDBFileModel) GetOffset() int64 {
	return s.Offset
}

func (s *SanDBFileModel) Write(entry sanface.EntryFace) error {
	buf, err := entry.Encode()
	if err != nil {
		fmt.Println("[Error] SanDB File Model user Func <entry.Encode> appear error:", err)
		return err
	}
	_, err = s.File.WriteAt(buf, s.GetOffset())
	if err != nil {
		fmt.Println("[Error] SanDB File Model Write entry appear error:", err)
		return err
	}
	s.Offset += entry.GetSize()
	return nil
}

func (s *SanDBFileModel) GetFile() *os.File {
	return s.File
}

func (s *SanDBFileModel) Read(offset int64) (sanface.EntryFace, error) {
	buf := make([]byte, conf.ConfigObj.EntryHeaderSize)
	_, err := s.File.ReadAt(buf, offset)
	if err != nil && err == io.EOF {
		fmt.Println("[Info] SanDBFile Read EOF And Exit:", err)
		return nil, err
	} else if err != nil {
		fmt.Println("[Error] SanDB Read File error:", err)
		return nil, err
	}
	entry, err := DecodeEntry(buf)
	if err != nil {
		fmt.Println("[Error] SanDB use func <Decode> error", err)
		return nil, err
	}
	offset += int64(conf.ConfigObj.EntryHeaderSize)
	if entry.GetKeySize() > 0 {
		keybuf := make([]byte, entry.GetKeySize())
		_, err = s.File.ReadAt(keybuf, offset)
		if err != nil {
			fmt.Println("[Error] SanDB Read File (get Key) error:", err)
			return nil, err
		}
		entry.SetKey(keybuf)
	}
	offset += int64(entry.GetKeySize())

	if entry.GetValSize() > 0 {
		valbuf := make([]byte, entry.GetValSize())
		_, err = s.File.ReadAt(valbuf, offset)
		if err != nil {
			fmt.Println("[Error] SanDB Read File (get Val) error:", err)
			return nil, err
		}
		entry.SetVal(valbuf)
	}
	return entry, nil
}

func NewSanDBFileModel(fileaddr string) (sanface.SanDBFileFace, error) {
	file, err := os.OpenFile(fileaddr, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("[Error] NewSanDBFileModel-NewSanDBFileModel user Func <os.OpenFile> appear error:", err)
		return nil, err
	}
	stat, err := os.Stat(fileaddr)
	if err != nil {
		fmt.Println("[Error] NewSanDBFileModel os.Stat appear error:", err)
		return nil, err
	}
	return &SanDBFileModel{
		file,
		stat.Size(),
	}, nil
}
