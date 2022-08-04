package sanmodel

import (
	"SanDB/sanface"
	"SanDB/tools"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SetRouteModel struct {
	ConnModel  *ConnModel
	setManager sanface.SetManagerFace
}

func (s *SetRouteModel) Sadd(trandata sanface.TranDataFace) error {
	kvlist := strings.Split(string(trandata.GetData()), " ")
	if len(kvlist) != 2 {
		_ = s.ConnModel.SendSyntaxMsg()
		return errors.New("[info] Syntax Error")
	}
	err := s.setManager.Sadd([]byte(kvlist[0]), []byte(kvlist[1]))
	if err != nil {
		fmt.Println("[Warning] Conn Slove Sadd user Func <sadd> appear Error:", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_ = s.ConnModel.SendSucessMsg()
	return nil
}
func (s *SetRouteModel) Scard(trandata sanface.TranDataFace) error {
	key := trandata.GetData()
	val := s.setManager.Scard(key)
	ans := NewTranDataModel([]byte(strconv.Itoa(val)), Suc)
	ansbuf, err := ans.Encode()
	if err != nil {
		fmt.Println("[Error] SetManager scard pack trandata error", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_, err = s.ConnModel.Conn.Write(ansbuf)
	if err != nil {
		fmt.Println("[Error] SetManager scard write into conn error", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	return nil
}

func (s *SetRouteModel) Smember(trandata sanface.TranDataFace) error {
	key := trandata.GetData()
	members, err := s.setManager.Smember(key)
	if err != nil {

	}
	msgbuf := tools.EncodeSetMember(members)
	td := NewTranDataModel(msgbuf, Suc)
	buf, err := td.Encode()
	if err != nil {
		fmt.Println("[Error] SetManager scard pack trandata error", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_, err = s.ConnModel.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] SetManager scard write into conn error", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	return nil
}

func (s *SetRouteModel) SIsmember(trandata sanface.TranDataFace) error {
	kvlist := strings.Split(string(trandata.GetData()), " ")
	if len(kvlist) != 2 {
		_ = s.ConnModel.SendSyntaxMsg()
		return errors.New("[info] Syntax Error")
	}
	flag := s.setManager.SIsMember([]byte(kvlist[0]), []byte(kvlist[1]))
	//if err != nil {
	//	fmt.Println("[Warning] Server use SIsmember error,", err)
	//	_ = s.ConnModel.SendErrMsg()
	//	return err
	//}
	var msgcontent string
	if flag == true {
		msgcontent = "true"
	} else {
		msgcontent = "false"
	}

	td := NewTranDataModel([]byte(msgcontent), Suc)
	buf, err := td.Encode()
	if err != nil {
		fmt.Println("[warning] SIsmember pack msg error")
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_, err = s.ConnModel.Conn.Write(buf)
	if err != nil {
		fmt.Println("[warning] Server Write msg error")
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	return nil
}

func (s *SetRouteModel) Spop(trandata sanface.TranDataFace) error {
	kvlist := strings.Split(string(trandata.GetData()), " ")
	if len(kvlist) != 2 {
		_ = s.ConnModel.SendSyntaxMsg()
		return errors.New("[info] Syntax Error")
	}
	err := s.setManager.Spop([]byte(kvlist[0]), []byte(kvlist[1]))
	if err != nil {
		fmt.Println("[Warning] Server use spop error,", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_ = s.ConnModel.SendSucessMsg()
	return nil
}
func (s *SetRouteModel) DelByKey(trandata sanface.TranDataFace) error {
	key := trandata.GetData()
	err := s.setManager.DelByKey(key)
	if err != nil {
		fmt.Println("[Warning] Server use DelByKey error,", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_ = s.ConnModel.SendSucessMsg()
	return nil
}

func (s *SetRouteModel) MergeFile(trandata sanface.TranDataFace) error {
	err := s.setManager.MergeFile()
	if err != nil {
		fmt.Println("[Warning] SetManager MergeFile error", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_ = s.ConnModel.SendSucessMsg()
	return nil
}

func (s *SetRouteModel) Clean() error {
	err := s.setManager.Clean()
	if err != nil {
		fmt.Println("[Warning] Setmanager clean error", err)
		_ = s.ConnModel.SendErrMsg()
		return err
	}
	_ = s.ConnModel.SendSucessMsg()
	return nil
}

func NewSetRouteModel(conn *ConnModel, sm sanface.SetManagerFace) *SetRouteModel {
	return &SetRouteModel{
		ConnModel:  conn,
		setManager: sm,
	}
}
