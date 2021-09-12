package container

import (
	"fmt"
)

type containerTask struct {
	taskType string
	options  interface{}
}

type containerWorker struct {
	runtime   ContainerRuntime
	dispacher ContainerDispatcher
	tasks     chan containerTask
	quit      chan bool
	number    int
}

type ContainerDispatcher struct {
	numWorkers int
	runtime    ContainerRuntime
	workerPool chan chan containerTask
	tasks      chan containerTask
	workers    []*containerWorker
}

func NewContainerDispatcher(numWorkers int, runtime ContainerRuntime) *ContainerDispatcher {
	return &ContainerDispatcher{
		numWorkers: numWorkers,
		runtime:    runtime,
		workerPool: make(chan chan containerTask, numWorkers),
		tasks:      make(chan containerTask),
	}
}

func (d *ContainerDispatcher) Start() {
	for i := 0; i < d.numWorkers; i++ {
		worker := newContainerWorker(d, d.runtime, i)
		d.workers = append(d.workers, worker)
		worker.listen()
	}
	go d.dispatch()
}

func (d *ContainerDispatcher) Stop() {
	for _, w := range d.workers {
		w.Stop()
	}
}

func (d *ContainerDispatcher) dispatch() {
	for {
		tk := <-d.tasks
		fmt.Printf("Dispatching %s\n", tk.taskType)
		go func() {
			w := <-d.workerPool
			w <- tk
		}()
	}
}

func newContainerWorker(dispacher *ContainerDispatcher, runtime ContainerRuntime, n int) *containerWorker {
	return &containerWorker{
		runtime:   runtime,
		dispacher: *dispacher,
		tasks:     make(chan (containerTask)),
		number:    n,
	}
}

func (c *containerWorker) listen() {
	go func() {
		var err error
	loop:
		for {
			c.dispacher.workerPool <- c.tasks // another worker is waiting for a task
			select {
			case t := <-c.tasks:
				switch t.taskType {
				case "buildImage":
					err = c.runtime.BuildImage(t.options.(BuildImageOptions))
				case "runContainer":
					err = c.runtime.RunContainer(t.options.(RunContainerOptions))
				case "pullImage":
					err = c.runtime.PullImage(t.options.(string))
				case "removeImage":
					err = c.runtime.RemoveImage(t.options.(string))
				}
			case <-c.quit:
				fmt.Printf("Worker: %d stopping ... \n", c.number+1)
				break loop
			}

			if err != nil {
				fmt.Println(err)
			}

		}
	}()
}

func (c *containerWorker) Stop() {
	go func() {
		c.quit <- true
	}()
}

func (d *ContainerDispatcher) BuildImage(options BuildImageOptions) error {
	d.tasks <- containerTask{taskType: "buildImage", options: options}
	return nil
}
func (d *ContainerDispatcher) PullImage(imageName string) error {
	d.tasks <- containerTask{taskType: "pullImage"}
	return nil
}
func (d *ContainerDispatcher) RunContainer(options RunContainerOptions) error {
	d.tasks <- containerTask{taskType: "runContainer", options: options}
	return nil
}
func (d *ContainerDispatcher) RemoveImage(imageName string) error {
	d.tasks <- containerTask{taskType: "removeImage"}
	return nil
}
