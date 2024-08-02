package utils

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const (
	workerBits uint8 = 1  // 每台机器(节点)的id位数 10位最大可以有2^10=1024个节点(0-1023)
	numberBits uint8 = 10 // 表示每个集群下的每个节点，1秒内可生成的id序号的二进制位数 即每秒最多可生成 2^22=4194303个唯一id(0-4194303)
	// 这里求最大值使用了位运算
	workerMax   int64 = ^(-1 << workerBits)     // 节点id的最大值，用于防止溢出
	numberMax   int64 = ^(-1 << numberBits)     // 每个节点，1秒内可生成的id序号最大值
	timeShift   uint8 = workerBits + numberBits // 时间戳向左的偏移量
	workerShift uint8 = numberBits              // 节点id向左的偏移量
	epoch       int64 = 1655794409              //这个是我在写epoch这个变量时的时间戳(秒)
)

// 定义一个woker工作节点所需要的基本参数
type Worker struct {
	mu        sync.Mutex // 添加互斥锁 确保并发安全
	timestamp int64      // 记录时间戳
	workerId  int64      // 该节点的id
	number    int64      // 当前毫秒已经生成的id序列号(从0开始累加) 1秒内最多生成4194304个id
}

var workerInst *Worker

func GetUuidInstance() *Worker {
	if workerInst == nil {
		workerInst, _ = NewWorker(1)
	}
	return workerInst
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("workId is invalidate")
	}
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetStringId() string {
	return strconv.FormatInt(w.GetId(), 10)
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().Unix()
	if now == w.timestamp {
		w.number++
		if w.number > numberMax {
			w.number = 0
			for now <= w.timestamp {
				now = time.Now().Unix()
			}
			w.timestamp = now
		}
	} else if now > w.timestamp {
		w.number = 0
		w.timestamp = now
	} else {
		for now < w.timestamp {
			now = time.Now().Unix()
		}
		w.number++
		if w.number > numberMax {
			w.number = 0
			for now <= w.timestamp {
				now = time.Now().Unix()
			}
			w.timestamp = now
		}
	}
	id := int64((now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number))
	return id
}
