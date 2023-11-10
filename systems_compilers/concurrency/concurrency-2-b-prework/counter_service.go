// Author: Patch Neranartkomol

package counterservice

import (
	"sync"
	"sync/atomic"
)

type CounterService interface {
	// Returns values in ascending order; it should be safe to call
	// getNext() concurrently from multiple goroutines without any
	// additional synchronization on the caller's side.
	getNext() uint64
}

type UnsynchronizedCounterService struct {
	/* Please implement this struct and its getNext method */
  Id uint64
}

// getNext() - This one can be UNSAFE
func (counter *UnsynchronizedCounterService) getNext() uint64 {
  counter.Id += 1
  return counter.Id
}

type AtomicCounterService struct {
  Id uint64
}

// getNext() with sync/atomic
func (counter *AtomicCounterService) getNext() uint64 {
  atomic.AddUint64(&counter.Id, 1)
  return counter.Id
}

type MutexCounterService struct {
  mu sync.Mutex
  Id uint64
}

// getNext() with sync/Mutex
func (counter *MutexCounterService) getNext() uint64 {
  counter.mu.Lock()
  counter.Id += 1
  counter.mu.Unlock()
  return counter.Id
}

type ChannelCounterService struct {
  inc chan bool
  count uint64
  done chan bool
}

// A constructor for ChannelCounterService
func newChannelCounterService() *ChannelCounterService {
	cs := ChannelCounterService{
    count: 0,
    inc: make(chan bool),
    done: make(chan bool),
  }
	return &cs
}

func(counter *ChannelCounterService) Increment(){
  counter.inc <- true
}
// getNext() with goroutines and channels
func (counter *ChannelCounterService) getNext() uint64 {

  counter.done <- true
  return <- counter.done
}
