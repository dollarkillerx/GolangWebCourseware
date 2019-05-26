package taskrunner

type Runner struct {
	Controller controllerChan // 控制信息
	Error controllerChan // 错误信息
	Data dataChan // 数据
	dataSize int // 数据大小
	longlived bool // 是否长期存货 tree不会回收资源
	Dispatcher fn // 生产者
	Executor fn // 消费者
}

func NewRunner(size int,longlived bool,dispatcher fn,executor fn) *Runner {
	return &Runner{
		Controller:make(chan string,1),// 我们这个需求是非阻塞的 所以用带buffer的chan
		Error:make(chan string,1),
		Data:make(chan interface{},size),
		longlived:longlived,
		Dispatcher:dispatcher,
		Executor:executor,
	}
}

func (r *Runner) startDispatch() {
	// 资源回收
	defer func() {
		if !r.longlived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	forloop:
	for  {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_TO_EXECUTE
				}
			}
			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil{
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:
			//如果执行这个 说明消费完了
			break forloop
		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
