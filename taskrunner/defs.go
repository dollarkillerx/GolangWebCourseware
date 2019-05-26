package taskrunner

const (
	READY_TO_DISPATCH = "d" // 开始生产 通知
	READY_TO_EXECUTE = "e" // 开始消费 通知
	CLOSE = "c" // 停止
)

type controllerChan chan string // 通知 通知处理

type dataChan chan interface{} // 存放处理数据

type fn func (dc dataChan) error // 处理方法