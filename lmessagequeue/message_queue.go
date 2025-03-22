package lmessagequeue

import (
	"sync"
	"time"
)

type Message[T any] struct {
	Content T
	DedupId string
}

type MessageQueue[T any] struct {
	ch       chan Message[T]
	dedupIds sync.Map
	timeout  time.Duration
	stopChan chan struct{}
	wg       sync.WaitGroup
}

func NewMessageQueue[T any](capacity int, timeout time.Duration) *MessageQueue[T] {
	mq := &MessageQueue[T]{
		ch:       make(chan Message[T], capacity),
		timeout:  timeout,
		stopChan: make(chan struct{}),
	}
	mq.wg.Add(1)
	go mq.cleanupExpiredDedupIds()
	return mq
}

func (mq *MessageQueue[T]) Send(msg Message[T]) bool {
	now := time.Now()
	if val, loaded := mq.dedupIds.Load(msg.DedupId); loaded {
		if timestamp, ok := val.(time.Time); ok && now.Sub(timestamp) < mq.timeout {
			return false
		}
	}
	mq.dedupIds.Store(msg.DedupId, now)
	select {
	case mq.ch <- msg:
		return true
	default:
		mq.dedupIds.Delete(msg.DedupId)
		return false
	}
}

func (mq *MessageQueue[T]) Receive() (Message[T], bool) {
	select {
	case msg, ok := <-mq.ch:
		return msg, ok
	default:
		var zero Message[T]
		return zero, false
	}
}

func (mq *MessageQueue[T]) Length() int {
	return len(mq.ch)
}

func (mq *MessageQueue[T]) cleanupExpiredDedupIds() {
	defer mq.wg.Done()
	if mq.timeout <= 0 {
		return
	}
	ticker := time.NewTicker(mq.timeout * 10)
	defer ticker.Stop()
	for {
		select {
		case <-mq.stopChan:
			return
		case <-ticker.C:
			now := time.Now()
			mq.dedupIds.Range(func(key, value interface{}) bool {
				if timestamp, ok := value.(time.Time); ok && now.Sub(timestamp) >= mq.timeout {
					mq.dedupIds.Delete(key)
				}
				return true
			})
		}
	}
}

func (mq *MessageQueue[T]) Close() {
	close(mq.stopChan)
	mq.wg.Wait()
	close(mq.ch)
}
