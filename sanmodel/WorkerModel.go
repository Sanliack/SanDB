package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"fmt"
	"time"
)

type WorkerModel struct {
	id        int
	MsgQueue  chan sanface.WorkerTranDataFace
	Closechan *chan int
}

func (w *WorkerModel) Start() {
	defer w.Stop()
	fmt.Printf("worker ID:%d is ready to Work....\n", w.id)
	for {
		if time.Now().Unix()%2 == 0 {
			fmt.Println("手动产生error")
			return
		}
		workermsg := <-w.MsgQueue
		err := workermsg.GetConn().SolveTranData(workermsg.GetData())
		if err != nil {
			return
		}
		fmt.Printf("worker%d 成功处理一条请求内容为%s:\n", w.id, string(workermsg.GetData().GetData()))
	}
}

func (w *WorkerModel) ReStart() {
	defer w.Stop()
	//w.MsgQueue = make(chan sanface.WorkerTranDataFace, conf.ConfigObj.WorkQueueLength)
	fmt.Printf("worker ID:%d Restart ....\n", w.id)
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
	//close(w.MsgQueue)
	*w.Closechan <- w.id
}

func (w *WorkerModel) AddMsg(data sanface.WorkerTranDataFace) {
	w.MsgQueue <- data
}

func NewWorkerModel(id int, closechan *chan int) sanface.WorkerFace {
	return &WorkerModel{
		id:        id,
		MsgQueue:  make(chan sanface.WorkerTranDataFace, conf.ConfigObj.WorkQueueLength),
		Closechan: closechan,
	}
}
