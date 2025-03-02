package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Task any

type Worker struct {
	ctx  context.Context
	chIn <-chan Task
}

type WorkerPool struct {
	ctx          context.Context
	mu           sync.Mutex
	workersLimit int
	tasksLimit   int
	chanIn       chan Task
	isClosed     atomic.Bool
}

func main() {
	fmt.Println("Worker Pool")
	ctx, cancel := context.WithCancel(context.Background())
	wp := NewWorkerPool(ctx, 5, 10)
	for i:= 0; i < 20; i++ {
		task := new(Task)
		wp.Add(task)
	}

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
	wp.Finish()
	task := new(Task)
	wp.Add(task)
}

func NewWorkerPool(ctx context.Context, workersLimit, tasksLimit int) *WorkerPool {
	chanIn := make(chan Task, tasksLimit)
	for i := 0; i < workersLimit; i++ {
		w := NewWorker(ctx, chanIn)
		w.Work()
	}

	return &WorkerPool{
		ctx:          ctx,
		workersLimit: workersLimit,
		tasksLimit:   tasksLimit,
		chanIn:       chanIn,
	}
}

func NewWorker(ctx context.Context, ch <-chan Task) *Worker{
	return &Worker{
		ctx:  ctx,
		chIn: ch,
	}
}

func (w *Worker) Work() {
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				fmt.Println("stop worker", w)
				return
			case t := <-w.chIn:
				time.Sleep(1 * time.Second)
				fmt.Println("task: ", t, "done")
			}
		}
	}()
}

func (wp *WorkerPool) Add(task Task) error {
	if wp.isClosed.Load() {
		fmt.Println("worker pool is closed")
		return errors.New("worker pool is closed")
	}
	wp.mu.Lock()
	defer wp.mu.Unlock()
	if len(wp.chanIn) == wp.tasksLimit {
		fmt.Println("too many tasks in queue")
		return errors.New("too many tasks in queue")
	} else {
		wp.chanIn <- task
	}

	return nil
}

func (wp *WorkerPool) Finish() {
	wp.isClosed.Store(true)
	close(wp.chanIn)
}
