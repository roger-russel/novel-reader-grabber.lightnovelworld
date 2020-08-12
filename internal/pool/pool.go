package pool

import (
	"fmt"
	"sync"
)

//Pool worker
type Pool struct {
	size     int
	workers  map[int]func()
	running  map[int]bool
	wg       sync.WaitGroup
	chDone   chan int
	tasks    int32 //The number of current tasks on pool
	finished int32
}

//New Pool
func New(size int) *Pool {
	p := Pool{
		size:    size,
		workers: make(map[int]func()),
		running: make(map[int]bool),
		chDone:  make(chan int, 1),
	}
	return &p
}

//Add Function into Pool
func (p *Pool) Add(f func()) {

	p.tasks++
	p.wg.Add(1)
	workerID := len(p.workers)
	p.workers[workerID] = func() {
		f()
		p.chDone <- workerID
	}

}

//Run Workers on Pool
func (p *Pool) Run() {
	go p.done()

	ln := len(p.workers)

	for workerID := 0; workerID < p.size && workerID < ln; workerID++ { // lock running go rotines
		p.running[workerID] = true
	}

	for workerID := 0; workerID < p.size && workerID < ln; workerID++ { // run it
		go p.workers[workerID]()
	}

	p.wg.Wait()

}

func (p *Pool) done() {

	for workerID := range p.chDone {
		delete(p.workers, workerID)
		delete(p.running, workerID)

		p.tasks--
		p.finished++
		p.wg.Done()

		if len(p.workers) > 0 {
			for next := range p.workers {
				if _, running := p.running[next]; !running {
					p.running[next] = true
					go p.workers[next]()
					break
				}
			}
		}
	}
}

//Status of pool
func (p *Pool) Status() string {
	return fmt.Sprintf("Getted: %v, Remaining: %v",
		p.finished,
		p.tasks,
	)
}
