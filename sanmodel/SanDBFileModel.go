package sanmodel

import (
	"fmt"
	"os"
)

type SanDBFileModel struct {
	File   *os.File
	Offset int
}

func NewSanDBFileModel(fileaddr string) *SanDBFileModel {
	file, err := os.OpenFile(fileaddr, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("[Error] NewSanDBFileModel appear error:", err)
		return nil
	}
	return &SanDBFileModel{
		file,
		0,
	}
}
