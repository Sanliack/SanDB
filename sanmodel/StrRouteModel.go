package sanmodel

import (
	"SanDB/sanface"
	"errors"
	"fmt"
	"strings"
)

type StrRouteModel struct {
	ConnModel   *ConnModel
	DataManager sanface.DataManagerFace
}

func (s *StrRouteModel) Get(trandata sanface.TranDataFace) error {
	key := trandata.GetData()
	val, err := s.DataManager.Get(key)
	remsg := NewTranDataModel(val, Suc)
	buf, err := remsg.Encode()
	if err != nil {
		fmt.Println("[Error] pack Remsg Error：", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	// cache===============================
	_, err = s.ConnModel.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] Conn Write Error：", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	return nil
}

func (s *StrRouteModel) Put(trandata sanface.TranDataFace) error {
	keyandval := strings.Split(string(trandata.GetData()), " ")
	if len(keyandval) != 2 {
		fmt.Println("[info] Accept Message syntax Error,pass")
		_ = s.ConnModel.SendSyntaxMsg()
		return errors.New("message syntax Error")
	}
	key := keyandval[0]
	val := keyandval[1]
	err := s.DataManager.Put([]byte(key), []byte(val))
	if err != nil {
		fmt.Println("[Warning] Conn Slove TranData user Func <conn.Put> appear Error:", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_ = s.ConnModel.SendSucessMsg()
	return nil
}

func (s *StrRouteModel) Del(trandata sanface.TranDataFace) error {
	key := trandata.GetData()
	err := s.DataManager.Del(key)
	if err != nil {
		fmt.Println("[Warning] Conn Slove TranData user Func <conn.Del> appear Error:", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_ = s.ConnModel.SendSucessMsg()
	return nil
}

func (s *StrRouteModel) Clean() error {
	err := s.DataManager.Clean()
	if err != nil {
		fmt.Println("[Warning] Conn Slove TranData user Func <conn.Cle> appear Error:", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_ = s.ConnModel.SendSucessMsg()
	return nil
}

func (s *StrRouteModel) Merge() error {
	err := s.DataManager.MergeFile()
	if err != nil {
		fmt.Println("[Error] Server User func <MergerFile> Error:", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_ = s.ConnModel.SendSucessMsg()
	return nil
	//case Set_Sadd:
}

func NewStrRouteModel(c *ConnModel, dm sanface.DataManagerFace) *StrRouteModel {
	return &StrRouteModel{
		ConnModel:   c,
		DataManager: dm,
	}
}
