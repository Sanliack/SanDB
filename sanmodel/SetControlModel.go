package sanmodel

import (
	"SanDB/sanface"
	"SanDB/tools"
	"errors"
	"fmt"
	"net"
	"strconv"
)

type SetControlModel struct {
	Conn *net.TCPConn
}

func (s *SetControlModel) Sadd(key []byte, val []byte) error {
	content := append(key, []byte(" ")...)
	content = append(content, val...)
	td := NewTranDataModel(content, Set_Add)
	buf, _ := td.Encode()
	_, err := s.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error]SetControlModel Send msg to Server appear error:", err)
		return err
	}
	remsg := make([]byte, 2048)
	n, err := s.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error]SetControlModel get msg from Server appear error:", err)
		return err
	}
	ent := DecodeTranData(remsg[:n])
	if ent.GetCommId() == Suc {
		return nil
	} else if ent.GetCommId() == Syntax {
		return errors.New("syntax Error")
	} else {
		return errors.New("something happen wrong")
	}
}
func (s *SetControlModel) Scard(key []byte) (int, error) {
	td := NewTranDataModel(key, Set_Card)
	buf, _ := td.Encode()
	_, err := s.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error]SetControlModel Send msg to Server appear error:", err)
		return -1, err
	}
	remsg := make([]byte, 2048)
	n, err := s.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error]SetControlModel get msg from Server appear error:", err)
		return -1, err
	}
	ent := DecodeTranData(remsg[:n])
	if ent.GetCommId() == Suc {
		nums, err := strconv.Atoi(string(ent.GetData()))
		if err != nil {
			fmt.Println("[Error] Server Get data is't a number")
			return -1, err
		}
		return nums, nil
	} else if ent.GetCommId() == Syntax {
		return -1, errors.New("syntax Error")
	} else {
		return -1, errors.New("something happen wrong")
	}
}

func (s *SetControlModel) Smember(key []byte) ([][]byte, error) {
	td := NewTranDataModel(key, Set_Member)
	buf, _ := td.Encode()
	_, err := s.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error]SetControlModel Send msg to Server appear error:", err)
		return nil, err
	}
	remsg := make([]byte, 2048)
	n, err := s.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error]SetControlModel get msg from Server appear error:", err)
		return nil, err
	}
	ent := DecodeTranData(remsg[:n])
	if ent.GetCommId() == Suc {
		member, err := tools.DecodeSetMember(ent.GetData())
		if err != nil {
			fmt.Println("[Info] tools.DecodeSetMember error，", err)
			return nil, err
		}
		return member, nil
	} else if ent.GetCommId() == Syntax {
		return nil, errors.New("syntax Error")
	} else {
		return nil, errors.New("something happen wrong")
	}
}

func (s *SetControlModel) Spop(key []byte, val []byte) error {
	content := append(key, []byte(" ")...)
	content = append(content, val...)
	td := NewTranDataModel(content, Set_Pop)
	buf, _ := td.Encode()
	_, err := s.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error]SetControlModel Send msg to Server appear error:", err)
		return err
	}
	remsg := make([]byte, 2048)
	n, err := s.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error]SetControlModel get msg from Server appear error:", err)
		return err
	}
	ent := DecodeTranData(remsg[:n])
	if ent.GetCommId() == Suc {
		return nil
	} else if ent.GetCommId() == Syntax {
		return errors.New("syntax Error")
	} else {
		return errors.New("something happen wrong")
	}
}
func (s *SetControlModel) SIsMember(key []byte, val []byte) (bool, error) {
	content := append(key, []byte(" ")...)
	content = append(content, val...)
	td := NewTranDataModel(content, Set_IsMember)
	buf, _ := td.Encode()
	_, err := s.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error]SetControlModel Send msg to Server appear error:", err)
		return false, err
	}
	remsg := make([]byte, 2048)
	n, err := s.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error]SetControlModel get msg from Server appear error:", err)
		return false, err
	}
	ent := DecodeTranData(remsg[:n])
	if ent.GetCommId() == Suc {
		data := ent.GetData()
		if string(data) == "true" {
			return true, nil
		} else if string(data) == "false" {
			return false, nil
		}
		fmt.Println("=============", string(ent.GetData()))
		return false, errors.New("[SIsMember] Accept Data is't true/false")

	} else if ent.GetCommId() == Syntax {
		return false, errors.New("syntax Error")
	} else {
		return false, errors.New("something happen wrong")
	}
}

func (s *SetControlModel) DelByKey(key []byte) error {
	td := NewTranDataModel(key, Set_DelByKey)
	buf, _ := td.Encode()
	_, err := s.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error]SetControlModel Send msg to Server appear error:", err)
		return err
	}
	remsg := make([]byte, 2048)
	n, err := s.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error]SetControlModel get msg from Server appear error:", err)
		return err
	}
	ent := DecodeTranData(remsg[:n])
	if ent.GetCommId() == Suc {
		return nil
	} else if ent.GetCommId() == Syntax {
		return errors.New("syntax Error")
	} else {
		return errors.New("something wrong happen")
	}
}

func (s *SetControlModel) MergeFile() error {
	td := NewTranDataModel(nil, Set_Merge)
	buf, _ := td.Encode()
	_, err := s.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error]SetControlModel Send set_merge msg to Server appear error:", err)
		return err
	}
	remsg := make([]byte, 2048)
	n, err := s.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error]SetControlModel get set_merge msg from Server appear error:", err)
		return err
	}
	ent := DecodeTranData(remsg[:n])
	if ent.GetCommId() == Suc {
		return nil
	} else if ent.GetCommId() == Syntax {
		return errors.New("syntax Error")
	} else {
		return errors.New("something wrong happen")
	}
}
func (s *SetControlModel) Clean() error {
	td := NewTranDataModel(nil, Set_Clean)
	buf, _ := td.Encode()
	_, err := s.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error]SetControlModel Send Str_Clean msg to Server appear error:", err)
		return err
	}
	remsg := make([]byte, 2048)
	n, err := s.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error]SetControlModel get Str_Clean msg from Server appear error:", err)
		return err
	}
	ent := DecodeTranData(remsg[:n])
	if ent.GetCommId() == Suc {
		return nil
	} else if ent.GetCommId() == Syntax {
		return errors.New("syntax Error")
	} else {
		return errors.New("something wrong happen")
	}
}

func NewSetControl(conn *net.TCPConn) sanface.SetControlFace {
	return &SetControlModel{
		Conn: conn,
	}
}
