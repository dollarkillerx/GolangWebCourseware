package taskrunner

import "time"

/**
	timer
	setup
	start{trigger->task->runner}
 */

type Worker struct {
	ticker *time.Ticker // 定时器
	runner *Runner
}

func NewWorker(interval time.Duration,r *Runner) *Worker {
	return &Worker{
		ticker:time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <- w.ticker.C://这个是阻塞的,当时间到了这里就放行了
			go w.runner.StartAll()
		}
	}
}

func Start() {
	// Start
	//r :NewWor   .....
}