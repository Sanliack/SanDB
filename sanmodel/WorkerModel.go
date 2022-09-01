package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"fmt"
)

type WorkerModel struct {
	id       int
	MsgQueue chan sanface.WorkerTranDataFace
}

func (w *WorkerModel) Start() {
	defer w.Stop()
	fmt.Printf("worker ID:%d is ready to Work....\n", w.id)
	for {
		workermsg := <-w.MsgQueue
		err := workermsg.GetConn().SolveTranData(workermsg.GetData())
		if err != nil {
			return
		}
		fmt.Printf("worker%d 成功处理一条请求内容为%s:\n", w.id, string(workermsg.GetData().GetData()))
	}
}

func (w *WorkerModel) Stop() {
	close(w.MsgQueue)
}

func (w *WorkerModel) AddMsg(data sanface.WorkerTranDataFace) {
	w.MsgQueue <- data
}

func NewWorkerModel(id int) sanface.WorkerFace {
	return &WorkerModel{
		id:       id,
		MsgQueue: make(chan sanface.WorkerTranDataFace, conf.ConfigObj.WorkQueueLength),
	}
}
