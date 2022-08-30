package server

type Worker interface {
	Add(func())
}

// SimpleWorker is a simple background worker that implements
// the Worker interface and runs all tasks in a go routine without
// a pool or que or limits. It's useful for simple or small applications
// with minimal/short background tasks
type SimpleWorker struct {
}

func NewSimpleWorker() *SimpleWorker {
	return &SimpleWorker{}
}

func (sw *SimpleWorker) Add(task func()) {
	go task()
}
