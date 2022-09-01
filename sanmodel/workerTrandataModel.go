package sanmodel

import "SanDB/sanface"

type WorkerTranData struct {
	Conn sanface.ConnFace
	Data sanface.TranDataFace
}

func (w *WorkerTranData) GetData() sanface.TranDataFace {
	return w.Data
}

func (w *WorkerTranData) GetConn() sanface.ConnFace {
	return w.Conn
}

func NewWorkerTranData(conn sanface.ConnFace, data sanface.TranDataFace) sanface.WorkerTranDataFace {
	return &WorkerTranData{
		Conn: conn,
		Data: data,
	}
}
