package runner

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal   // 中断信号
	complete  chan error       // 结束时状态符
	timeout   <-chan time.Time // 超时时间
	tasks     []func(int)      // 任务列表
}

var ErrTimeout = errors.New("Error for timeout")
var ErrInterrupt = errors.New("Error for Interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1), // 保证 goroutine 准备好去接受中断信号
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// 添加任务
func (r *Runner) Add(task ...func(int)) {
	r.tasks = append(r.tasks, task...)
}

func (r *Runner) Start() error {

	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		log.Printf("receive err %+v\n", err)
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	// check for an interrupt signal from the OS.
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	// singal when an interrupt event is sent.
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
